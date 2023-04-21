const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List trusted access role bindings.
 *
 * @summary List trusted access role bindings.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-03-02-preview/examples/TrustedAccessRoleBindings_List.json
 */
async function listTrustedAccessRoleBindings() {
  const subscriptionId = process.env["CONTAINERSERVICE_SUBSCRIPTION_ID"] || "subid1";
  const resourceGroupName = process.env["CONTAINERSERVICE_RESOURCE_GROUP"] || "rg1";
  const resourceName = "clustername1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.trustedAccessRoleBindings.list(resourceGroupName, resourceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
