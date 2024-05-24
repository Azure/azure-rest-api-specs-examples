const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to VM actions on DbNode of VM Cluster by the provided filter
 *
 * @summary VM actions on DbNode of VM Cluster by the provided filter
 * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/dbNodes_action.json
 */
async function vmActionsOnDbNodesOfVMCluster() {
  const subscriptionId =
    process.env["ORACLEDATABASE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["ORACLEDATABASE_RESOURCE_GROUP"] || "rg000";
  const cloudvmclustername = "cluster1";
  const dbnodeocid = "ocid1....aaaaaa";
  const body = { action: "Start" };
  const credential = new DefaultAzureCredential();
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const result = await client.dbNodes.beginActionAndWait(
    resourceGroupName,
    cloudvmclustername,
    dbnodeocid,
    body,
  );
  console.log(result);
}
