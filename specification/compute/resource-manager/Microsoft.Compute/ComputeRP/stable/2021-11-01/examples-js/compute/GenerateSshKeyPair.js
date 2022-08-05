const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function generateAnSshKeyPair() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const sshPublicKeyName = "mySshPublicKeyName";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.sshPublicKeys.generateKeyPair(resourceGroupName, sshPublicKeyName);
  console.log(result);
}

generateAnSshKeyPair().catch(console.error);
