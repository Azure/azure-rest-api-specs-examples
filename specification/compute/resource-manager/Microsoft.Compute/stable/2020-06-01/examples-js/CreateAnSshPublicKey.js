const { ComputeManagementClient } = require("@azure/arm-compute-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new SSH public key resource.
 *
 * @summary Creates a new SSH public key resource.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2020-06-01/examples/CreateAnSshPublicKey.json
 */
async function createANewSshPublicKeyResource() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const sshPublicKeyName = "mySshPublicKeyName";
  const parameters = {
    location: "westus",
    publicKey: "{ssh-rsa public key}",
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.sshPublicKeys.create(resourceGroupName, sshPublicKeyName, parameters);
  console.log(result);
}
