const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a KPack build result.
 *
 * @summary Get a KPack build result.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-09-01-preview/examples/BuildService_GetBuildResult.json
 */
async function buildServiceGetBuildResult() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const buildServiceName = "default";
  const buildName = "mybuild";
  const buildResultName = "123";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.buildServiceOperations.getBuildResult(
    resourceGroupName,
    serviceName,
    buildServiceName,
    buildName,
    buildResultName
  );
  console.log(result);
}

buildServiceGetBuildResult().catch(console.error);
