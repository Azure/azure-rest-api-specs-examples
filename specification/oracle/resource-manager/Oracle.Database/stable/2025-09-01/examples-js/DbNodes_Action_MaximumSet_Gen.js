const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to vM actions on DbNode of VM Cluster by the provided filter
 *
 * @summary vM actions on DbNode of VM Cluster by the provided filter
 * x-ms-original-file: 2025-09-01/DbNodes_Action_MaximumSet_Gen.json
 */
async function vmActionsOnDbNodesOfVMClusterGeneratedByMaximumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const result = await client.dbNodes.action("rgopenapi", "cloudvmcluster1", "abciderewdidsereq", {
    action: "Start",
  });
  console.log(result);
}
