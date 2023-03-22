const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a scope map with the specified parameters.
 *
 * @summary Updates a scope map with the specified parameters.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2022-12-01/examples/ScopeMapUpdate.json
 */
async function scopeMapUpdate() {
  const subscriptionId =
    process.env["CONTAINERREGISTRY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONTAINERREGISTRY_RESOURCE_GROUP"] || "myResourceGroup";
  const registryName = "myRegistry";
  const scopeMapName = "myScopeMap";
  const scopeMapUpdateParameters = {
    description: "Developer Scopes",
    actions: ["repositories/myrepository/contentWrite", "repositories/myrepository/contentRead"],
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.scopeMaps.beginUpdateAndWait(
    resourceGroupName,
    registryName,
    scopeMapName,
    scopeMapUpdateParameters
  );
  console.log(result);
}
