package registry

import (
	"context"
	"encoding/csv"
	"io"
	"log"
	"os"
)

type CSVToMutationLoader interface {
	LineToMutation(context.Context, chan error, []string) chan string
}

// LoadFromCSV loads data from a CSV into a .rdf file with triples in  RDF N-Quad format.
// input file should be CSV and out put generated will be a .rdf file.
func LoadFromCSV(logger *log.Logger, inputFile, outputFile string, mutationLoader CSVToMutationLoader) error {
	input, err := os.Open(inputFile)
	if err != nil {
		return err
	}

	output, err := os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}

	reader := csv.NewReader(input)
	ctx := context.Background()

	var (
		recordChan = make(chan []string)
		done       = make(chan bool, 1)
		errorChan  = make(chan error)
	)

	defer func() {
		output.Close()
		close(recordChan)
		close(done)
		close(errorChan)
	}()
	for {
		go func() {
			record, err := reader.Read()
			if err == io.EOF {
				done <- true
			}
			if err != nil {
				errorChan <- err
				return
			}
			recordChan <- record
		}()

		select {
		case record := <-recordChan:
			go func() {
				mutChan := mutationLoader.LineToMutation(ctx, errorChan, record)
				mut := <-mutChan
				_, err := output.Write([]byte(mut))
				if err != nil {
					errorChan <- err
				}
				//logger.Printf("wrote sub query %v", mut)
			}()
		case err := <-errorChan:
			// if there is an error in writing one line we print the error and continue
			logger.Fatalf("error reading csv and writing to rdf %v", err)
		case <-done:
			output.Close()
			close(recordChan)
			close(done)
			close(errorChan)
			os.Exit(0)
		}
	}

}
