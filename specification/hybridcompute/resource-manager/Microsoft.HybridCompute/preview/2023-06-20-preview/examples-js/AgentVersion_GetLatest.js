const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an Agent Version along with the download link currently present.
 *
 * @summary Gets an Agent Version along with the download link currently present.
 * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/AgentVersion_GetLatest.json
 */
async function getAgentVersions() {
  const osType = "myOsType";
  const version = "1.27";
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential);
  const result = await client.agentVersionOperations.get(osType, version);
  console.log(result);
}
