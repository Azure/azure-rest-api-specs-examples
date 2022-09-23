const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete an SSH public key.
 *
 * @summary Delete an SSH public key.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/sshPublicKeyExamples/SshPublicKeys_Delete_MaximumSet_Gen.json
 */
async function sshPublicKeysDeleteMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const sshPublicKeyName = "aaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.sshPublicKeys.delete(resourceGroupName, sshPublicKeyName);
  console.log(result);
}

sshPublicKeysDeleteMaximumSetGen().catch(console.error);
