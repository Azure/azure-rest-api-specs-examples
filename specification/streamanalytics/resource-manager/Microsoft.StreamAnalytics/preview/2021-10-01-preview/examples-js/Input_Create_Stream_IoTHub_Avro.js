const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates an input or replaces an already existing input under an existing streaming job.
 *
 * @summary Creates an input or replaces an already existing input under an existing streaming job.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Input_Create_Stream_IoTHub_Avro.json
 */
async function createAStreamIoTHubInputWithAvroSerialization() {
  const subscriptionId =
    process.env["STREAMANALYTICS_SUBSCRIPTION_ID"] || "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = process.env["STREAMANALYTICS_RESOURCE_GROUP"] || "sjrg3467";
  const jobName = "sj9742";
  const inputName = "input7970";
  const input = {
    properties: {
      type: "Stream",
      datasource: {
        type: "Microsoft.Devices/IotHubs",
        consumerGroupName: "sdkconsumergroup",
        endpoint: "messages/events",
        iotHubNamespace: "iothub",
        sharedAccessPolicyKey: "sharedAccessPolicyKey=",
        sharedAccessPolicyName: "owner",
      },
      serialization: { type: "Avro" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.inputs.createOrReplace(resourceGroupName, jobName, inputName, input);
  console.log(result);
}
