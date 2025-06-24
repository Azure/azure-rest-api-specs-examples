const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets all Extension publishers based on the location
 *
 * @summary Gets all Extension publishers based on the location
 * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/extension/ExtensionPublisher_List.json
 */
async function getAListOfExtensionPublishers() {
  const location = "EastUS";
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential);
  const resArray = new Array();
  for await (const item of client.extensionPublisherOperations.list(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}
