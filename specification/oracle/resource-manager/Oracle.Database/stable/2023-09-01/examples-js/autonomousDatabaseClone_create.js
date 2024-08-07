const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a AutonomousDatabase
 *
 * @summary Create a AutonomousDatabase
 * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/autonomousDatabaseClone_create.json
 */
async function createCloneAutonomousDatabase() {
  const subscriptionId =
    process.env["ORACLEDATABASE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["ORACLEDATABASE_RESOURCE_GROUP"] || "rg000";
  const autonomousdatabasename = "databasedb1";
  const resource = {
    location: "eastus",
    properties: {
      adminPassword: "********",
      characterSet: "AL32UTF8",
      cloneType: "Full",
      computeCount: 2,
      computeModel: "ECPU",
      dataBaseType: "Clone",
      dataStorageSizeInTbs: 1,
      displayName: "example_autonomous_databasedb1_clone",
      ncharacterSet: "AL16UTF16",
      sourceId:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/autonomousDatabases/databasedb1",
      subnetId:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1",
      vnetId:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Microsoft.Network/virtualNetworks/vnet1",
    },
    tags: { tagK1: "tagV1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const result = await client.autonomousDatabases.beginCreateOrUpdateAndWait(
    resourceGroupName,
    autonomousdatabasename,
    resource,
  );
  console.log(result);
}
