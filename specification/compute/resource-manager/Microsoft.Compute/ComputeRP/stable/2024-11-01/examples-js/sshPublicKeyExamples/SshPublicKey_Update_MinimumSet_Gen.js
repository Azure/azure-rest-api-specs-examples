const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Updates a new SSH public key resource.
 *
 * @summary Updates a new SSH public key resource.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/sshPublicKeyExamples/SshPublicKey_Update_MinimumSet_Gen.json
 */
async function sshPublicKeyUpdateMinimumSetGen() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "rgcompute";
  const sshPublicKeyName = "aaaaaaaaaaa";
  const parameters = {};
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.sshPublicKeys.update(resourceGroupName, sshPublicKeyName, parameters);
  console.log(result);
}
