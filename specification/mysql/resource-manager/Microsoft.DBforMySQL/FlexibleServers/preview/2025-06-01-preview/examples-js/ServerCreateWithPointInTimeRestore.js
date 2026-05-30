const { MySQLManagementFlexibleServerClient } = require("@azure/arm-mysql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new server or updates an existing server. The update action will overwrite the existing server.
 *
 * @summary creates a new server or updates an existing server. The update action will overwrite the existing server.
 * x-ms-original-file: 2025-06-01-preview/ServerCreateWithPointInTimeRestore.json
 */
async function createAServerAsAPointInTimeRestore() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new MySQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.servers.create("TargetResourceGroup", "targetserver", {
    location: "SoutheastAsia",
    createMode: "PointInTimeRestore",
    restorePointInTime: new Date("2021-06-24T00:00:37.467Z"),
    sourceServerResourceId:
      "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/SourceResourceGroup/providers/Microsoft.DBforMySQL/flexibleServers/sourceserver",
    sku: { name: "Standard_D14_v2", tier: "GeneralPurpose" },
    tags: { num: "1" },
  });
  console.log(result);
}
