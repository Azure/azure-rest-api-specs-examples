const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List KPack builds.
 *
 * @summary List KPack builds.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/BuildService_ListBuilds.json
 */
async function buildServiceListBuilds() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const buildServiceName = "default";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.buildServiceOperations.listBuilds(
    resourceGroupName,
    serviceName,
    buildServiceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

buildServiceListBuilds().catch(console.error);
