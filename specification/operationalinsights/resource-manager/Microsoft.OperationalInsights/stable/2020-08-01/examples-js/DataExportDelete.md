```javascript
const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified data export in a given workspace..
 *
 * @summary Deletes the specified data export in a given workspace..
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/DataExportDelete.json
 */
async function dataExportDelete() {
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = "RgTest1";
  const workspaceName = "DeWnTest1234";
  const dataExportName = "export1";
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const result = await client.dataExports.delete(resourceGroupName, workspaceName, dataExportName);
  console.log(result);
}

dataExportDelete().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-operationalinsights_8.0.1/sdk/operationalinsights/arm-operationalinsights/README.md) on how to add the SDK to your project and authenticate.
