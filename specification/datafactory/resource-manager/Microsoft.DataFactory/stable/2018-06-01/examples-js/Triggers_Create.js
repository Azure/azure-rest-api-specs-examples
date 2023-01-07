const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a trigger.
 *
 * @summary Creates or updates a trigger.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_Create.json
 */
async function triggersCreate() {
  const subscriptionId =
    process.env["DATAFACTORY_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = process.env["DATAFACTORY_RESOURCE_GROUP"] || "exampleResourceGroup";
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
