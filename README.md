# realty-tool

This is a backend service that uses dgraph as its storage, data loaded to dgraph is generated by csvload.go, 
the resulting rdf file is used to generate a bulk load data set by running:
    
    `dgraph bulk -f output.rdf.gz -s schemaProperty.dql.gz`

Once bulk files are generated it's copied to alpha pods created using `node-three-bulk.yaml` deployment file and then 
the init-container used to populate alpha container.
