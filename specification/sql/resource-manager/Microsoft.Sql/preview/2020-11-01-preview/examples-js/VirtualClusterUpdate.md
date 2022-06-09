```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a virtual cluster.
 *
 * @summary Updates a virtual cluster.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/VirtualClusterUpdate.json
 */
async function updateVirtualClusterWithTags() {
  const subscriptionId = "20d7082a-0fc7-4468-82bd-542694d5042b";
  const resourceGroupName = "testrg";
  const virtualClusterName = "vc-subnet1-f769ed71-b3ad-491a-a9d5-26eeceaa6be2";
  const parameters = {
    maintenanceConfigurationId:
      "/subscriptions/ab0e51c0-83c0-4380-8ae9-025516df392f/resourceGroups/Federation/providers/Microsoft.Maintenance/maintenanceConfigurations/MiPolicy1",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.virtualClusters.beginUpdateAndWait(
    resourceGroupName,
    virtualClusterName,
    parameters
  );
  console.log(result);
}

updateVirtualClusterWithTags().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.
