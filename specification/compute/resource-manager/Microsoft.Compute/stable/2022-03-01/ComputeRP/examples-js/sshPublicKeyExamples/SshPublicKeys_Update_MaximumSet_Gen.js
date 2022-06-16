const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a new SSH public key resource.
 *
 * @summary Updates a new SSH public key resource.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/sshPublicKeyExamples/SshPublicKeys_Update_MaximumSet_Gen.json
 */
async function sshPublicKeysUpdateMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const sshPublicKeyName = "aaaaaaaaaaaa";
  const parameters = {
    publicKey: "{ssh-rsa public key}",
    tags: { key2854: "a" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.sshPublicKeys.update(resourceGroupName, sshPublicKeyName, parameters);
  console.log(result);
}

sshPublicKeysUpdateMaximumSetGen().catch(console.error);
