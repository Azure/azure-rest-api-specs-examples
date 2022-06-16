const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a pipeline.
 *
 * @summary Creates or updates a pipeline.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Pipelines_Create.json
 */
async function pipelinesCreate() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const pipelineName = "examplePipeline";
  const pipeline = {
    activities: [
      {
        name: "ExampleForeachActivity",
        type: "ForEach",
        activities: [
          {
            name: "ExampleCopyActivity",
            type: "Copy",
            dataIntegrationUnits: 32,
            inputs: [
              {
                type: "DatasetReference",
                parameters: {
                  myFileName: "examplecontainer.csv",
                  myFolderPath: "examplecontainer",
                },
                referenceName: "exampleDataset",
              },
            ],
            outputs: [
              {
                type: "DatasetReference",
                parameters: {
                  myFileName: { type: "Expression", value: "@item()" },
                  myFolderPath: "examplecontainer",
                },
                referenceName: "exampleDataset",
              },
            ],
            sink: { type: "BlobSink" },
            source: { type: "BlobSource" },
          },
        ],
        isSequential: true,
        items: {
          type: "Expression",
          value: "@pipeline().parameters.OutputBlobNameList",
        },
      },
    ],
    parameters: {
      jobId: { type: "String" },
      outputBlobNameList: { type: "Array" },
    },
    policy: { elapsedTimeMetric: { duration: "0.00:10:00" } },
    runDimensions: {
      jobId: { type: "Expression", value: "@pipeline().parameters.JobId" },
    },
    variables: { testVariableArray: { type: "Array" } },
  };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.pipelines.createOrUpdate(
    resourceGroupName,
    factoryName,
    pipelineName,
    pipeline
  );
  console.log(result);
}

pipelinesCreate().catch(console.error);
