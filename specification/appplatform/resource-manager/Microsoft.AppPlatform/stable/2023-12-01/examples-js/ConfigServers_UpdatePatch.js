const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the config server.
 *
 * @summary Update the config server.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/ConfigServers_UpdatePatch.json
 */
async function configServersUpdatePatch() {
  const subscriptionId =
    process.env["APPPLATFORM_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPPLATFORM_RESOURCE_GROUP"] || "myResourceGroup";
  const serviceName = "myservice";
  const configServerResource = {
    properties: {
      configServer: {
        gitProperty: {
          label: "master",
          searchPaths: ["/"],
          uri: "https://github.com/fake-user/fake-repository.git",
        },
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.configServers.beginUpdatePatchAndWait(
    resourceGroupName,
    serviceName,
    configServerResource,
  );
  console.log(result);
}
