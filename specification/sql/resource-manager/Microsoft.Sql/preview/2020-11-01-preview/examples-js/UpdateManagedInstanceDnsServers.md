```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Synchronizes the DNS server settings used by the managed instances inside the given virtual cluster.
 *
 * @summary Synchronizes the DNS server settings used by the managed instances inside the given virtual cluster.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/UpdateManagedInstanceDnsServers.json
 */
async function synchronizesTheDnsServerSettingsUsedByTheManagedInstancesInsideTheGivenVirtualCluster() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "sqlcrudtest-7398";
  const virtualClusterName = "sqlcrudtest-4645";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.virtualClusters.updateDnsServers(
    resourceGroupName,
    virtualClusterName
  );
  console.log(result);
}

synchronizesTheDnsServerSettingsUsedByTheManagedInstancesInsideTheGivenVirtualCluster().catch(
  console.error
);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.
