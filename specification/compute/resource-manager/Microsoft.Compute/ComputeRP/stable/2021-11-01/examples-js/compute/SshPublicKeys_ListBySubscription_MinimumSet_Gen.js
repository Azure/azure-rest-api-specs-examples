const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the SSH public keys in the subscription. Use the nextLink property in the response to get the next page of SSH public keys.
 *
 * @summary Lists all of the SSH public keys in the subscription. Use the nextLink property in the response to get the next page of SSH public keys.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/SshPublicKeys_ListBySubscription_MinimumSet_Gen.json
 */
async function sshPublicKeysListBySubscriptionMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.sshPublicKeys.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

sshPublicKeysListBySubscriptionMinimumSetGen().catch(console.error);
