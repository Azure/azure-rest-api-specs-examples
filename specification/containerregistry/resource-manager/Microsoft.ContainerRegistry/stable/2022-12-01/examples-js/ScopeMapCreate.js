const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a scope map for a container registry with the specified parameters.
 *
 * @summary Creates a scope map for a container registry with the specified parameters.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2022-12-01/examples/ScopeMapCreate.json
 */
async function scopeMapCreate() {
  const subscriptionId =
    process.env["CONTAINERREGISTRY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONTAINERREGISTRY_RESOURCE_GROUP"] || "myResourceGroup";
  const registryName = "myRegistry";
  const scopeMapName = "myScopeMap";
  const scopeMapCreateParameters = {
    description: "Developer Scopes",
    actions: ["repositories/myrepository/contentWrite", "repositories/myrepository/delete"],
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.scopeMaps.beginCreateAndWait(
    resourceGroupName,
    registryName,
    scopeMapName,
    scopeMapCreateParameters
  );
  console.log(result);
}
