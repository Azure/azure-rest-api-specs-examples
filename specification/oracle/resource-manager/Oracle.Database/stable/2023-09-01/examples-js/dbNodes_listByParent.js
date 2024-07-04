const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List DbNode resources by CloudVmCluster
 *
 * @summary List DbNode resources by CloudVmCluster
 * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/dbNodes_listByParent.json
 */
async function listDbNodesByVMCluster() {
  const subscriptionId =
    process.env["ORACLEDATABASE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["ORACLEDATABASE_RESOURCE_GROUP"] || "rg000";
  const cloudvmclustername = "cluster1";
  const credential = new DefaultAzureCredential();
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.dbNodes.listByCloudVmCluster(
    resourceGroupName,
    cloudvmclustername,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
