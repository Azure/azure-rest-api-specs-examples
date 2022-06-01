Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_18.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete an SSH public key.
 *
 * @summary Delete an SSH public key.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/sshPublicKeyExamples/SshPublicKeys_Delete_MaximumSet_Gen.json
 */
async function sshPublicKeysDeleteMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const sshPublicKeyName = "aaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.sshPublicKeys.delete(resourceGroupName, sshPublicKeyName);
  console.log(result);
}

sshPublicKeysDeleteMaximumSetGen().catch(console.error);
```
