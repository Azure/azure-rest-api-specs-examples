const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a trusted access role binding.
 *
 * @summary Get a trusted access role binding.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-08-02-preview/examples/TrustedAccessRoleBindings_Get.json
 */
async function getATrustedAccessRoleBinding() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const trustedAccessRoleBindingName = "binding1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.trustedAccessRoleBindings.get(
    resourceGroupName,
    resourceName,
    trustedAccessRoleBindingName
  );
  console.log(result);
}

getATrustedAccessRoleBinding().catch(console.error);
