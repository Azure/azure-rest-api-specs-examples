const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function createANewSshPublicKeyResource() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
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

createANewSshPublicKeyResource().catch(console.error);
