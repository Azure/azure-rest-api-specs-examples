const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the config server.
 *
 * @summary Update the config server.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-09-01-preview/examples/ConfigServers_UpdatePut.json
 */
async function configServersUpdatePut() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
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
  const result = await client.configServers.beginUpdatePutAndWait(
    resourceGroupName,
    serviceName,
    configServerResource
  );
  console.log(result);
}

configServersUpdatePut().catch(console.error);
