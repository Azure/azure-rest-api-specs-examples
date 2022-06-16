const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an instance pool.
 *
 * @summary Creates or updates an instance pool.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateInstancePoolMin.json
 */
async function createAnInstancePoolWithMinProperties() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "group1";
  const instancePoolName = "testIP";
  const parameters = {
    licenseType: "LicenseIncluded",
    location: "japaneast",
    sku: { name: "GP_Gen5", family: "Gen5", tier: "GeneralPurpose" },
    subnetId:
      "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Network/virtualNetworks/myvnet/subnets/mysubnet1",
    vCores: 8,
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.instancePools.beginCreateOrUpdateAndWait(
    resourceGroupName,
    instancePoolName,
    parameters
  );
  console.log(result);
}

createAnInstancePoolWithMinProperties().catch(console.error);
