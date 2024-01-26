const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update container registry resource.
 *
 * @summary Create or update container registry resource.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/ContainerRegistries_CreateOrUpdate.json
 */
async function containerRegistriesCreateOrUpdate() {
  const subscriptionId =
    process.env["APPPLATFORM_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPPLATFORM_RESOURCE_GROUP"] || "myResourceGroup";
  const serviceName = "my-service";
  const containerRegistryName = "my-container-registry";
  const containerRegistryResource = {
    properties: {
      credentials: {
        type: "BasicAuth",
        password: "myPassword",
        server: "myServer",
        username: "myUsername",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.containerRegistries.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceName,
    containerRegistryName,
    containerRegistryResource,
  );
  console.log(result);
}
