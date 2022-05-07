Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-datafactory_10.4.0/sdk/datafactory/arm-datafactory/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

async function pipelinesUpdate() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const pipelineName = "examplePipeline";
  const pipeline = {
    description: "Example description",
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
    parameters: { outputBlobNameList: { type: "Array" } },
    policy: { elapsedTimeMetric: { duration: "0.00:10:00" } },
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

pipelinesUpdate().catch(console.error);
```
