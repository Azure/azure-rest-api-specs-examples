const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a KPack build result log download URL.
 *
 * @summary Get a KPack build result log download URL.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/BuildService_GetBuildResultLog.json
 */
async function buildServiceGetBuildResultLog() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const buildServiceName = "default";
  const buildName = "mybuild";
  const buildResultName = "123";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.buildServiceOperations.getBuildResultLog(
    resourceGroupName,
    serviceName,
    buildServiceName,
    buildName,
    buildResultName
  );
  console.log(result);
}

buildServiceGetBuildResultLog().catch(console.error);
