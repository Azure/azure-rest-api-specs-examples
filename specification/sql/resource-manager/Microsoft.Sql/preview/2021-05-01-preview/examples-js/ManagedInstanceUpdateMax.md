Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a managed instance.
 *
 * @summary Updates a managed instance.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ManagedInstanceUpdateMax.json
 */
async function updateManagedInstanceWithAllProperties() {
  const subscriptionId = "20D7082A-0FC7-4468-82BD-542694D5042B";
  const resourceGroupName = "testrg";
  const managedInstanceName = "testinstance";
  const parameters = {
    administratorLogin: "dummylogin",
    administratorLoginPassword: "PLACEHOLDER",
    collation: "SQL_Latin1_General_CP1_CI_AS",
    licenseType: "BasePrice",
    maintenanceConfigurationId:
      "/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_JapanEast_MI_1",
    minimalTlsVersion: "1.2",
    proxyOverride: "Redirect",
    publicDataEndpointEnabled: false,
    requestedBackupStorageRedundancy: "Geo",
    sku: { name: "GP_Gen4", capacity: 8, tier: "GeneralPurpose" },
    storageSizeInGB: 448,
    tags: { tagKey1: "TagValue1" },
    vCores: 8,
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedInstances.beginUpdateAndWait(
    resourceGroupName,
    managedInstanceName,
    parameters
  );
  console.log(result);
}

updateManagedInstanceWithAllProperties().catch(console.error);
```
