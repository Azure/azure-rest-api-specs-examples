const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an instance pool.
 *
 * @summary creates or updates an instance pool.
 * x-ms-original-file: 2025-01-01/CreateOrUpdateInstancePoolMin.json
 */
async function createAnInstancePoolWithMinProperties() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.instancePools.createOrUpdate("group1", "testIP", {
    location: "japaneast",
    licenseType: "LicenseIncluded",
    subnetId:
      "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Network/virtualNetworks/myvnet/subnets/mysubnet1",
    vCores: 8,
    sku: { name: "GP_Gen5", family: "Gen5", tier: "GeneralPurpose" },
  });
  console.log(result);
}
