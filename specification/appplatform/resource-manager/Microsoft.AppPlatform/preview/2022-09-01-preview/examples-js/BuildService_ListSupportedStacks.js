const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get all supported stacks.
 *
 * @summary Get all supported stacks.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-09-01-preview/examples/BuildService_ListSupportedStacks.json
 */
async function buildServiceListSupportedStacks() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const buildServiceName = "default";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.buildServiceOperations.listSupportedStacks(
    resourceGroupName,
    serviceName,
    buildServiceName
  );
  console.log(result);
}

buildServiceListSupportedStacks().catch(console.error);
