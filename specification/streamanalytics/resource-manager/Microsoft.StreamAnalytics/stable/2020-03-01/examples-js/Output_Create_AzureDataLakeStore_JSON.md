Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-streamanalytics_4.0.1/sdk/streamanalytics/arm-streamanalytics/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates an output or replaces an already existing output under an existing streaming job.
 *
 * @summary Creates an output or replaces an already existing output under an existing streaming job.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Output_Create_AzureDataLakeStore_JSON.json
 */
async function createAnAzureDataLakeStoreOutputWithJsonSerialization() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = "sjrg6912";
  const jobName = "sj3310";
  const outputName = "output5195";
  const output = {
    datasource: {
      type: "Microsoft.DataLake/Accounts",
      accountName: "someaccount",
      dateFormat: "yyyy/MM/dd",
      filePathPrefix: "{date}/{time}",
      refreshToken: "someRefreshToken==",
      tenantId: "cea4e98b-c798-49e7-8c40-4a2b3beb47dd",
      timeFormat: "HH",
      tokenUserDisplayName: "Bob Smith",
      tokenUserPrincipalName: "bobsmith@contoso.com",
    },
    serialization: { type: "Json", format: "Array", encoding: "UTF8" },
  };
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.outputs.createOrReplace(
    resourceGroupName,
    jobName,
    outputName,
    output
  );
  console.log(result);
}

createAnAzureDataLakeStoreOutputWithJsonSerialization().catch(console.error);
```
