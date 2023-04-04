const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Operation to update an exiting Binding.
 *
 * @summary Operation to update an exiting Binding.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2023-03-01-preview/examples/Bindings_Update.json
 */
async function bindingsUpdate() {
  const subscriptionId =
    process.env["APPPLATFORM_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPPLATFORM_RESOURCE_GROUP"] || "myResourceGroup";
  const serviceName = "myservice";
  const appName = "myapp";
  const bindingName = "mybinding";
  const bindingResource = {
    properties: {
      bindingParameters: { apiType: "SQL", databaseName: "db1" },
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
