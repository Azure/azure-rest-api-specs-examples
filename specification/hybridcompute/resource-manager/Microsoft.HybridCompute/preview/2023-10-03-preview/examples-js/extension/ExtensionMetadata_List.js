const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all Extension versions based on location, publisher, extensionType
 *
 * @summary Gets all Extension versions based on location, publisher, extensionType
 * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-10-03-preview/examples/extension/ExtensionMetadata_List.json
 */
async function getAListOfExtensions() {
  const subscriptionId =
    process.env["HYBRIDCOMPUTE_SUBSCRIPTION_ID"] || "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const location = "EastUS";
  const publisher = "microsoft.azure.monitor";
  const extensionType = "azuremonitorlinuxagent";
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.extensionMetadata.list(location, publisher, extensionType)) {
    resArray.push(item);
  }
  console.log(resArray);
}
