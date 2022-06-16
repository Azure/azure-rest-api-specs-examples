const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Operation to delete a Binding.
 *
 * @summary Operation to delete a Binding.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Bindings_Delete.json
 */
async function bindingsDelete() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const appName = "myapp";
  const bindingName = "mybinding";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.bindings.beginDeleteAndWait(
    resourceGroupName,
    serviceName,
    appName,
    bindingName
  );
  console.log(result);
}

bindingsDelete().catch(console.error);
