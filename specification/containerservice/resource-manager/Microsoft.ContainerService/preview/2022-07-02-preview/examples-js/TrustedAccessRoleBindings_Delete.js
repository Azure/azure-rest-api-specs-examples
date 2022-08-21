const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a trusted access role binding.
 *
 * @summary Delete a trusted access role binding.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-07-02-preview/examples/TrustedAccessRoleBindings_Delete.json
 */
async function deleteATrustedAccessRoleBinding() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const trustedAccessRoleBindingName = "binding1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.trustedAccessRoleBindings.delete(
    resourceGroupName,
    resourceName,
    trustedAccessRoleBindingName
  );
  console.log(result);
}

deleteATrustedAccessRoleBinding().catch(console.error);
