const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the SSH public keys in the specified resource group. Use the nextLink property in the response to get the next page of SSH public keys.
 *
 * @summary Lists all of the SSH public keys in the specified resource group. Use the nextLink property in the response to get the next page of SSH public keys.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/sshPublicKeyExamples/SshPublicKey_ListByResourceGroup_MinimumSet_Gen.json
 */
async function sshPublicKeyListByResourceGroupMinimumSetGen() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "rgcompute";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.sshPublicKeys.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
