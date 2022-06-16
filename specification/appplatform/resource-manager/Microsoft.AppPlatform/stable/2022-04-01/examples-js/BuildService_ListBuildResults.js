const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List KPack build results.
 *
 * @summary List KPack build results.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/BuildService_ListBuildResults.json
 */
async function buildServiceListBuildResults() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const buildServiceName = "default";
  const buildName = "mybuild";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.buildServiceOperations.listBuildResults(
    resourceGroupName,
    serviceName,
    buildServiceName,
    buildName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

buildServiceListBuildResults().catch(console.error);
