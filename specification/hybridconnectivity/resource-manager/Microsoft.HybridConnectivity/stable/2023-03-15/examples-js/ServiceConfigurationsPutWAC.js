const { HybridConnectivityManagementAPI } = require("@azure/arm-hybridconnectivity");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a service in serviceConfiguration for the endpoint resource.
 *
 * @summary Create or update a service in serviceConfiguration for the endpoint resource.
 * x-ms-original-file: specification/hybridconnectivity/resource-manager/Microsoft.HybridConnectivity/stable/2023-03-15/examples/ServiceConfigurationsPutWAC.json
 */
async function serviceConfigurationsPutWac() {
  const resourceUri =
    "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine/providers/Microsoft.HybridConnectivity/endpoints/default";
  const endpointName = "default";
  const serviceConfigurationName = "WAC";
  const serviceConfigurationResource = {
    port: 6516,
    serviceName: "WAC",
  };
  const credential = new DefaultAzureCredential();
  const client = new HybridConnectivityManagementAPI(credential);
  const result = await client.serviceConfigurations.createOrupdate(
    resourceUri,
    endpointName,
    serviceConfigurationName,
    serviceConfigurationResource
  );
  console.log(result);
}
