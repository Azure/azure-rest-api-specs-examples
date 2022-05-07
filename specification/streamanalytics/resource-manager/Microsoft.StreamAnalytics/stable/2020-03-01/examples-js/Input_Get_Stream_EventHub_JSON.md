Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-streamanalytics_4.0.1/sdk/streamanalytics/arm-streamanalytics/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets details about the specified input.
 *
 * @summary Gets details about the specified input.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Input_Get_Stream_EventHub_JSON.json
 */
async function getAStreamEventHubInputWithJsonSerialization() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = "sjrg3139";
  const jobName = "sj197";
  const inputName = "input7425";
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.inputs.get(resourceGroupName, jobName, inputName);
  console.log(result);
}

getAStreamEventHubInputWithJsonSerialization().catch(console.error);
```
