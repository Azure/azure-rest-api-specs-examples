const { SiteRecoveryManagementClient } = require("@azure/arm-recoveryservices-siterecovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the details of an ASR replication protected item.
 *
 * @summary gets the details of an ASR replication protected item.
 * x-ms-original-file: 2026-11-01/ReplicationProtectedItems_Get_A2ACvm.json
 */
async function getsTheDetailsOfAnA2AConfidentialVMWithCMKDiskEncryption() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "42195872-7e70-4f8a-837f-84b28ecbb78b";
  const client = new SiteRecoveryManagementClient(credential, subscriptionId);
  const result = await client.replicationProtectedItems.get(
    "priyanprg",
    "priyanponeboxvault",
    "CentralUSCanSite",
    "CentralUSCancloud",
    "a2aCvmVm",
  );
  console.log(result);
}
