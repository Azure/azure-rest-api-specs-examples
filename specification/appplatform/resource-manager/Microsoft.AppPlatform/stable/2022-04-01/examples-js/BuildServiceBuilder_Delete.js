const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a KPack builder.
 *
 * @summary Delete a KPack builder.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/BuildServiceBuilder_Delete.json
 */
async function buildServiceBuilderDelete() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const buildServiceName = "default";
  const builderName = "mybuilder";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.buildServiceBuilder.beginDeleteAndWait(
    resourceGroupName,
    serviceName,
    buildServiceName,
    builderName
  );
  console.log(result);
}

buildServiceBuilderDelete().catch(console.error);
