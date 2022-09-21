const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Operation to update an exiting Binding.
 *
 * @summary Operation to update an exiting Binding.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-09-01-preview/examples/Bindings_Update.json
 */
async function bindingsUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const appName = "myapp";
  const bindingName = "mybinding";
  const bindingResource = {
    properties: {
      bindingParameters: { apiType: { string: "SQL" }, databaseName: { string: "db1" } },
      createdAt: undefined,
      generatedProperties: undefined,
      key: "xxxx",
      updatedAt: undefined,
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.bindings.beginUpdateAndWait(
    resourceGroupName,
    serviceName,
    appName,
    bindingName,
    bindingResource
  );
  console.log(result);
}

bindingsUpdate().catch(console.error);
