const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Add VMs to the VM Cluster
 *
 * @summary Add VMs to the VM Cluster
 * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/vmClusters_addVms.json
 */
async function addVMSToVMCluster() {
  const subscriptionId =
    process.env["ORACLEDATABASE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["ORACLEDATABASE_RESOURCE_GROUP"] || "rg000";
  const cloudvmclustername = "cluster1";
  const body = { dbServers: ["ocid1..aaaa", "ocid1..aaaaaa"] };
  const credential = new DefaultAzureCredential();
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const result = await client.cloudVmClusters.beginAddVmsAndWait(
    resourceGroupName,
    cloudvmclustername,
    body,
  );
  console.log(result);
}
