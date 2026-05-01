const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a trigger.
 *
 * @summary creates or updates a trigger.
 * x-ms-original-file: 2018-06-01/Triggers_Update.json
 */
async function triggersUpdate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789012";
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.triggers.createOrUpdate(
    "exampleResourceGroup",
    "exampleFactoryName",
    "exampleTrigger",
    {
      properties: {
        type: "ScheduleTrigger",
        description: "Example description",
        pipelines: [
          {
            parameters: { OutputBlobNameList: ["exampleoutput.csv"] },
            pipelineReference: { type: "PipelineReference", referenceName: "examplePipeline" },
          },
        ],
        recurrence: {
          endTime: new Date("2018-06-16T00:55:14.905167Z"),
          frequency: "Minute",
          interval: 4,
          startTime: new Date("2018-06-16T00:39:14.905167Z"),
          timeZone: "UTC",
        },
      },
    },
  );
  console.log(result);
}
