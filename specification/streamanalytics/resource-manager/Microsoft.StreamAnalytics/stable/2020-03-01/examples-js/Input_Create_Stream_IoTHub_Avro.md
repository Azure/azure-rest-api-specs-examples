Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-streamanalytics_4.0.1/sdk/streamanalytics/arm-streamanalytics/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates an input or replaces an already existing input under an existing streaming job.
 *
 * @summary Creates an input or replaces an already existing input under an existing streaming job.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Input_Create_Stream_IoTHub_Avro.json
 */
async function createAStreamIoTHubInputWithAvroSerialization() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = "sjrg3467";
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

createAStreamIoTHubInputWithAvroSerialization().catch(console.error);
```
