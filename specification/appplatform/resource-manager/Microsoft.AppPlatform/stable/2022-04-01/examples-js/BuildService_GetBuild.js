const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a KPack build.
 *
 * @summary Get a KPack build.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/BuildService_GetBuild.json
 */
async function buildServiceGetBuild() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const buildServiceName = "default";
  const buildName = "mybuild";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.buildServiceOperations.getBuild(
    resourceGroupName,
    serviceName,
    buildServiceName,
    buildName
  );
  console.log(result);
}

buildServiceGetBuild().catch(console.error);
