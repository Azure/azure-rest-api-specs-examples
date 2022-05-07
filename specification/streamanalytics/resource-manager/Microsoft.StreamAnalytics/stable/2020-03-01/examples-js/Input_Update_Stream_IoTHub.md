Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-streamanalytics_4.0.1/sdk/streamanalytics/arm-streamanalytics/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing input under an existing streaming job. This can be used to partially update (ie. update one or two properties) an input without affecting the rest the job or input definition.
 *
 * @summary Updates an existing input under an existing streaming job. This can be used to partially update (ie. update one or two properties) an input without affecting the rest the job or input definition.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Input_Update_Stream_IoTHub.json
 */
async function updateAStreamIoTHubInput() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = "sjrg3467";
  const jobName = "sj9742";
  const inputName = "input7970";
  const input = {
    properties: {
      type: "Stream",
      datasource: {
        type: "Microsoft.Devices/IotHubs",
        endpoint: "messages/operationsMonitoringEvents",
      },
      serialization: { type: "Csv", encoding: "UTF8", fieldDelimiter: "|" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.inputs.update(resourceGroupName, jobName, inputName, input);
  console.log(result);
}

updateAStreamIoTHubInput().catch(console.error);
```
