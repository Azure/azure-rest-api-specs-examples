const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a VirtualNetworkAddress
 *
 * @summary Delete a VirtualNetworkAddress
 * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/virtualNetworkAddresses_delete.json
 */
async function deleteVirtualNetworkAddress() {
  const subscriptionId =
    process.env["ORACLEDATABASE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["ORACLEDATABASE_RESOURCE_GROUP"] || "rg000";
  const cloudvmclustername = "cluster1";
  const virtualnetworkaddressname = "hostname1";
  const credential = new DefaultAzureCredential();
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkAddresses.beginDeleteAndWait(
    resourceGroupName,
    cloudvmclustername,
    virtualnetworkaddressname,
  );
  console.log(result);
}
