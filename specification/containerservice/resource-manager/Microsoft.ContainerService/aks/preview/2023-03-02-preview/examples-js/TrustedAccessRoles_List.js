const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List supported trusted access roles.
 *
 * @summary List supported trusted access roles.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-03-02-preview/examples/TrustedAccessRoles_List.json
 */
async function listTrustedAccessRoles() {
  const subscriptionId = process.env["CONTAINERSERVICE_SUBSCRIPTION_ID"] || "subid1";
  const location = "westus2";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.trustedAccessRoles.list(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}
