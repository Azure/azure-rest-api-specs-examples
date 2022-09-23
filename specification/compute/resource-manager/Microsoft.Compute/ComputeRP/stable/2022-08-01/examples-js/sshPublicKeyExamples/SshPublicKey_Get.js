const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves information about an SSH public key.
 *
 * @summary Retrieves information about an SSH public key.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/sshPublicKeyExamples/SshPublicKey_Get.json
 */
async function getAnSshPublicKey() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "myResourceGroup";
  const sshPublicKeyName = "mySshPublicKeyName";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.sshPublicKeys.get(resourceGroupName, sshPublicKeyName);
  console.log(result);
}

getAnSshPublicKey().catch(console.error);
