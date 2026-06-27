const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a managed instance.
 *
 * @summary creates or updates a managed instance.
 * x-ms-original-file: 2025-01-01/ManagedInstanceCreateMin.json
 */
async function createManagedInstanceWithMinimalProperties() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "20D7082A-0FC7-4468-82BD-542694D5042B";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedInstances.createOrUpdate("testrg", "testinstance", {
    location: "Japan East",
    administratorLogin: "dummylogin",
    administratorLoginPassword: "PLACEHOLDER",
    licenseType: "LicenseIncluded",
    storageSizeInGB: 1024,
    subnetId:
      "/subscriptions/20D7082A-0FC7-4468-82BD-542694D5042B/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1",
    vCores: 8,
    sku: { name: "GP_Gen5", tier: "GeneralPurpose" },
  });
  console.log(result);
}
