const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a managed instance.
 *
 * @summary Creates or updates a managed instance.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ManagedInstanceCreateMax.json
 */
async function createManagedInstanceWithAllProperties() {
  const subscriptionId = "20D7082A-0FC7-4468-82BD-542694D5042B";
  const resourceGroupName = "testrg";
  const managedInstanceName = "testinstance";
  const parameters = {
    administratorLogin: "dummylogin",
    administratorLoginPassword: "PLACEHOLDER",
    administrators: {
      azureADOnlyAuthentication: true,
      login: "bob@contoso.com",
      principalType: "User",
      sid: "00000011-1111-2222-2222-123456789111",
      tenantId: "00000011-1111-2222-2222-123456789111",
    },
    collation: "SQL_Latin1_General_CP1_CI_AS",
    dnsZonePartner:
      "/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/resourceGroups/testrg/providers/Microsoft.Sql/managedInstances/testinstance",
    instancePoolId:
      "/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/resourceGroups/testrg/providers/Microsoft.Sql/instancePools/pool1",
    licenseType: "LicenseIncluded",
    location: "Japan East",
    maintenanceConfigurationId:
      "/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_JapanEast_MI_1",
    minimalTlsVersion: "1.2",
    proxyOverride: "Redirect",
    publicDataEndpointEnabled: false,
    requestedBackupStorageRedundancy: "Geo",
    servicePrincipal: { type: "SystemAssigned" },
    sku: { name: "GP_Gen5", tier: "GeneralPurpose" },
    storageSizeInGB: 1024,
    subnetId:
      "/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1",
    tags: { tagKey1: "TagValue1" },
    timezoneId: "UTC",
    vCores: 8,
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedInstances.beginCreateOrUpdateAndWait(
    resourceGroupName,
    managedInstanceName,
    parameters
  );
  console.log(result);
}

createManagedInstanceWithAllProperties().catch(console.error);
