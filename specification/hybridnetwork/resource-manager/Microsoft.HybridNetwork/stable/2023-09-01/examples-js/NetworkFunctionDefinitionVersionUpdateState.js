const { HybridNetworkManagementClient } = require("@azure/arm-hybridnetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update network function definition version state.
 *
 * @summary Update network function definition version state.
 * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/NetworkFunctionDefinitionVersionUpdateState.json
 */
async function updateNetworkFunctionDefinitionVersionState() {
  const subscriptionId = process.env["HYBRIDNETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HYBRIDNETWORK_RESOURCE_GROUP"] || "rg";
  const publisherName = "TestPublisher";
  const networkFunctionDefinitionGroupName = "TestSkuGroup";
  const networkFunctionDefinitionVersionName = "1.0.0";
  const parameters = {
    versionState: "Active",
  };
  const credential = new DefaultAzureCredential();
  const client = new HybridNetworkManagementClient(credential, subscriptionId);
  const result = await client.networkFunctionDefinitionVersions.beginUpdateStateAndWait(
    resourceGroupName,
    publisherName,
    networkFunctionDefinitionGroupName,
    networkFunctionDefinitionVersionName,
    parameters,
  );
  console.log(result);
}
