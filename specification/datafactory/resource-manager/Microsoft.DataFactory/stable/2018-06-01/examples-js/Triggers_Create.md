Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-datafactory_10.5.0/sdk/datafactory/arm-datafactory/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

async function triggersCreate() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const triggerName = "exampleTrigger";
  const trigger = {
    properties: {
      type: "ScheduleTrigger",
      pipelines: [
        {
          parameters: { outputBlobNameList: ["exampleoutput.csv"] },
          pipelineReference: {
            type: "PipelineReference",
            referenceName: "examplePipeline",
          },
        },
      ],
      recurrence: {
        endTime: new Date("2018-06-16T00:55:13.8441801Z"),
        frequency: "Minute",
        interval: 4,
        startTime: new Date("2018-06-16T00:39:13.8441801Z"),
        timeZone: "UTC",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.triggers.createOrUpdate(
    resourceGroupName,
    factoryName,
    triggerName,
    trigger
  );
  console.log(result);
}

triggersCreate().catch(console.error);
```
