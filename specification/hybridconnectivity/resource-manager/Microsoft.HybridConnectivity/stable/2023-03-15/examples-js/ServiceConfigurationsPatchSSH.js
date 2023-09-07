const { HybridConnectivityManagementAPI } = require("@azure/arm-hybridconnectivity");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the service details in the service configurations of the target resource.
 *
 * @summary Update the service details in the service configurations of the target resource.
 * x-ms-original-file: specification/hybridconnectivity/resource-manager/Microsoft.HybridConnectivity/stable/2023-03-15/examples/ServiceConfigurationsPatchSSH.json
 */
async function serviceConfigurationsPatchSsh() {
  const resourceUri =
    "subscriptions/f5bcc1d9-23af-4ae9-aca1-041d0f593a63/resourceGroups/hybridRG/providers/Microsoft.HybridCompute/machines/testMachine/providers/Microsoft.HybridConnectivity/endpoints/default";
  const endpointName = "default";
  const serviceConfigurationName = "SSH";
  const serviceConfigurationResource = {
    port: 22,
  };
  const credential = new DefaultAzureCredential();
  const client = new HybridConnectivityManagementAPI(credential);
  const result = await client.serviceConfigurations.update(
    resourceUri,
    endpointName,
    serviceConfigurationName,
    serviceConfigurationResource
  );
  console.log(result);
}
