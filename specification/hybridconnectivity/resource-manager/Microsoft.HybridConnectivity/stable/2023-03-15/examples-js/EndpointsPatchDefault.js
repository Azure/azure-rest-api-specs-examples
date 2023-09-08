const { HybridConnectivityManagementAPI } = require("@azure/arm-hybridconnectivity");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the endpoint to the target resource.
 *
 * @summary Update the endpoint to the target resource.
 * x-ms-original-file: specification/hybridconnectivity/resource-manager/Microsoft.HybridConnectivity/stable/2023-03-15/examples/EndpointsPatchDefault.json
 */
async function hybridConnectivityEndpointsPatchDefault() {
  const resourceUri =
    "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine";
  const endpointName = "default";
  const endpointResource = {
    properties: { type: "default" },
  };
  const credential = new DefaultAzureCredential();
  const client = new HybridConnectivityManagementAPI(credential);
  const result = await client.endpoints.update(resourceUri, endpointName, endpointResource);
  console.log(result);
}
