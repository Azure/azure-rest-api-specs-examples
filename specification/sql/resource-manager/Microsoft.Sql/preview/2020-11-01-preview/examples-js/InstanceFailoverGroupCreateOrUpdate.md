Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a failover group.
 *
 * @summary Creates or updates a failover group.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/InstanceFailoverGroupCreateOrUpdate.json
 */
async function createFailoverGroup() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default";
  const locationName = "Japan East";
  const failoverGroupName = "failover-group-test-3";
  const parameters = {
    managedInstancePairs: [
      {
        partnerManagedInstanceId:
          "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/failover-group-secondary-mngdInstance",
        primaryManagedInstanceId:
          "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/managedInstances/failover-group-primary-mngdInstance",
      },
    ],
    partnerRegions: [{ location: "Japan West" }],
    readOnlyEndpoint: { failoverPolicy: "Disabled" },
    readWriteEndpoint: {
      failoverPolicy: "Automatic",
      failoverWithDataLossGracePeriodMinutes: 480,
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.instanceFailoverGroups.beginCreateOrUpdateAndWait(
    resourceGroupName,
    locationName,
    failoverGroupName,
    parameters
  );
  console.log(result);
}

createFailoverGroup().catch(console.error);
```
