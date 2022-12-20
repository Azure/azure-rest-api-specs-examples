const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the list of outbound network dependencies for a given Azure-SSIS integration runtime.
 *
 * @summary Gets the list of outbound network dependencies for a given Azure-SSIS integration runtime.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_ListOutboundNetworkDependenciesEndpoints.json
 */
async function getOutboundNetworkDependencyEndpoints() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "ade9c2b6-c160-4305-9bb9-80342f6c1ae2";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "exampleResourceGroup";
  const workspaceName = "exampleWorkspace";
  const integrationRuntimeName = "exampleIntegrationRuntime";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.integrationRuntimes.listOutboundNetworkDependenciesEndpoints(
    resourceGroupName,
    workspaceName,
    integrationRuntimeName
  );
  console.log(result);
}
