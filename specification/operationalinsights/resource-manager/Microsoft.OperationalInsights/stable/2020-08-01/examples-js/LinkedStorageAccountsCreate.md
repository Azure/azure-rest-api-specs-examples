```javascript
const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or Update a link relation between current workspace and a group of storage accounts of a specific data source type.
 *
 * @summary Create or Update a link relation between current workspace and a group of storage accounts of a specific data source type.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/LinkedStorageAccountsCreate.json
 */
async function linkedStorageAccountsCreate() {
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = "mms-eus";
  const workspaceName = "testLinkStorageAccountsWS";
  const dataSourceType = "CustomLogs";
  const parameters = {
    storageAccountIds: [
      "/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/mms-eus/providers/Microsoft.Storage/storageAccounts/testStorageA",
      "/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/mms-eus/providers/Microsoft.Storage/storageAccounts/testStorageB",
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const result = await client.linkedStorageAccounts.createOrUpdate(
    resourceGroupName,
    workspaceName,
    dataSourceType,
    parameters
  );
  console.log(result);
}

linkedStorageAccountsCreate().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-operationalinsights_8.0.1/sdk/operationalinsights/arm-operationalinsights/README.md) on how to add the SDK to your project and authenticate.
