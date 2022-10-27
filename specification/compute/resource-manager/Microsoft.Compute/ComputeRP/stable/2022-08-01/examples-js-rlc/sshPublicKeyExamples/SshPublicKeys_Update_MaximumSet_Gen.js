const createComputeManagementClient = require("@azure-rest/arm-compute").default;
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to Updates a new SSH public key resource.
 *
 * @summary Updates a new SSH public key resource.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/sshPublicKeyExamples/SshPublicKeys_Update_MaximumSet_Gen.json
 */
async function sshPublicKeysUpdateMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const client = createComputeManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "rgcompute";
  const sshPublicKeyName = "aaaaaaaaaaaa";
  const options = {
    body: {
      properties: { publicKey: "{ssh-rsa public key}" },
      tags: { key2854: "a" },
    },
    queryParameters: { "api-version": "2022-08-01" },
  };
  const result = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/sshPublicKeys/{sshPublicKeyName}",
      subscriptionId,
      resourceGroupName,
      sshPublicKeyName
    )
    .patch(options);
  console.log(result);
}

sshPublicKeysUpdateMaximumSetGen().catch(console.error);
