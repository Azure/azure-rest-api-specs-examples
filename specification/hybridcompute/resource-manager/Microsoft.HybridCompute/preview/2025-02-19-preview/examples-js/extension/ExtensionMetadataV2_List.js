const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets all Extension versions based on location, publisher, extensionType
 *
 * @summary Gets all Extension versions based on location, publisher, extensionType
 * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/extension/ExtensionMetadataV2_List.json
 */
async function getAListOfExtensionMetadata() {
  const location = "EastUS";
  const publisher = "microsoft.azure.monitor";
  const extensionType = "azuremonitorlinuxagent";
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential);
  const resArray = new Array();
  for await (const item of client.extensionMetadataV2.list(location, publisher, extensionType)) {
    resArray.push(item);
  }
  console.log(resArray);
}
