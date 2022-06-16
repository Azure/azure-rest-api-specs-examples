const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function sshPublicKeysUpdateMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const sshPublicKeyName = "aaaaaaaaaaa";
  const parameters = {};
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.sshPublicKeys.update(resourceGroupName, sshPublicKeyName, parameters);
  console.log(result);
}

sshPublicKeysUpdateMinimumSetGen().catch(console.error);
