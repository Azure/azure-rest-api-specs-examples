const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List trusted access role bindings.
 *
 * @summary List trusted access role bindings.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-07-02-preview/examples/TrustedAccessRoleBindings_List.json
 */
async function listTrustedAccessRoleBindings() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.trustedAccessRoleBindings.list(resourceGroupName, resourceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listTrustedAccessRoleBindings().catch(console.error);
