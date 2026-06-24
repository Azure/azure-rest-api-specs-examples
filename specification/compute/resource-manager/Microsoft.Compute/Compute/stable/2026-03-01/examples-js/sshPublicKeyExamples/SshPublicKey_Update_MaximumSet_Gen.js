const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates a new SSH public key resource.
 *
 * @summary updates a new SSH public key resource.
 * x-ms-original-file: 2026-03-01/sshPublicKeyExamples/SshPublicKey_Update_MaximumSet_Gen.json
 */
async function sshPublicKeyUpdateMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.sshPublicKeys.update("rgcompute", "aaaaaaaaaaaa", {
    publicKey: "{ssh-rsa public key}",
    tags: { key2854: "a" },
  });
  console.log(result);
}
