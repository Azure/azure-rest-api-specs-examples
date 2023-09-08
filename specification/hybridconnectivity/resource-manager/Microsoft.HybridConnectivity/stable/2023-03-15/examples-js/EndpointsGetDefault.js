const { HybridConnectivityManagementAPI } = require("@azure/arm-hybridconnectivity");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the endpoint to the resource.
 *
 * @summary Gets the endpoint to the resource.
 * x-ms-original-file: specification/hybridconnectivity/resource-manager/Microsoft.HybridConnectivity/stable/2023-03-15/examples/EndpointsGetDefault.json
 */
async function hybridConnectivityEndpointsGetDefault() {
  const resourceUri =
    "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine";
  const endpointName = "default";
  const credential = new DefaultAzureCredential();
  const client = new HybridConnectivityManagementAPI(credential);
  const result = await client.endpoints.get(resourceUri, endpointName);
  console.log(result);
}
