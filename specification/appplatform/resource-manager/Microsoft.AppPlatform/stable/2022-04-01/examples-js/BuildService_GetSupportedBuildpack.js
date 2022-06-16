const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the supported buildpack resource.
 *
 * @summary Get the supported buildpack resource.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/BuildService_GetSupportedBuildpack.json
 */
async function buildServiceGetSupportedBuildpack() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const buildServiceName = "default";
  const buildpackName = "tanzu-buildpacks-java-azure";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.buildServiceOperations.getSupportedBuildpack(
    resourceGroupName,
    serviceName,
    buildServiceName,
    buildpackName
  );
  console.log(result);
}

buildServiceGetSupportedBuildpack().catch(console.error);
