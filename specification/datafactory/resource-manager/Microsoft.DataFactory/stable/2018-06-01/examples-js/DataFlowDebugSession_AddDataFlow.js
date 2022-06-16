const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Add a data flow into debug session.
 *
 * @summary Add a data flow into debug session.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlowDebugSession_AddDataFlow.json
 */
async function dataFlowDebugSessionAddDataFlow() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const request = {
    dataFlow: {
      name: "dataflow1",
      properties: {
        type: "MappingDataFlow",
        script:
          "\n\nsource(output(\n		Column_1 as string\n	),\n	allowSchemaDrift: true,\n	validateSchema: false) ~> source1",
        sinks: [],
        sources: [
          {
            name: "source1",
            dataset: {
              type: "DatasetReference",
              referenceName: "DelimitedText2",
            },
          },
        ],
        transformations: [],
      },
    },
    datasets: [
      {
        name: "dataset1",
        properties: {
          type: "DelimitedText",
          schema: [{ type: "String" }],
          annotations: [],
          columnDelimiter: ",",
          escapeChar: "\\",
          firstRowAsHeader: true,
          linkedServiceName: {
            type: "LinkedServiceReference",
            referenceName: "linkedService5",
          },
          location: {
            type: "AzureBlobStorageLocation",
            container: "dataflow-sample-data",
            fileName: "Ansiencoding.csv",
          },
          quoteChar: '"',
        },
      },
    ],
    debugSettings: {
      datasetParameters: { Movies: { path: "abc" }, Output: { time: "def" } },
      parameters: { sourcePath: "Toy" },
      sourceSettings: [
        { rowLimit: 1000, sourceName: "source1" },
        { rowLimit: 222, sourceName: "source2" },
      ],
    },
    linkedServices: [
      {
        name: "linkedService1",
        properties: {
          type: "AzureBlobStorage",
          annotations: [],
          connectionString:
            "DefaultEndpointsProtocol=https;AccountName=<storageName>;EndpointSuffix=core.windows.net;",
          encryptedCredential: "<credential>",
        },
      },
    ],
    sessionId: "f06ed247-9d07-49b2-b05e-2cb4a2fc871e",
  };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.dataFlowDebugSession.addDataFlow(
    resourceGroupName,
    factoryName,
    request
  );
  console.log(result);
}

dataFlowDebugSessionAddDataFlow().catch(console.error);
