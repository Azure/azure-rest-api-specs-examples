const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates a new SSH public key resource.
 *
 * @summary updates a new SSH public key resource.
 * x-ms-original-file: 2025-11-01/sshPublicKeyExamples/SshPublicKey_Update_MinimumSet_Gen.json
 */
async function sshPublicKeyUpdateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.sshPublicKeys.update("rgcompute", "aaaaaaaaaaa", {});
  console.log(result);
}
