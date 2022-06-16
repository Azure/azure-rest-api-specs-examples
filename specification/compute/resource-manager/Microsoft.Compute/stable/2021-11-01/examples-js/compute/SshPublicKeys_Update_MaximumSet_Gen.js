const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

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
