const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create the default Application Configuration Service or update the existing Application Configuration Service.
 *
 * @summary Create the default Application Configuration Service or update the existing Application Configuration Service.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/ConfigurationServices_CreateOrUpdate.json
 */
async function configurationServicesCreateOrUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const configurationServiceName = "default";
  const configurationServiceResource = {
    properties: {
      settings: {
        gitProperty: {
          repositories: [
            {
              name: "fake",
              label: "master",
              patterns: ["app/dev"],
              uri: "https://github.com/fake-user/fake-repository",
            },
          ],
        },
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.configurationServices.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceName,
    configurationServiceName,
    configurationServiceResource
  );
  console.log(result);
}

configurationServicesCreateOrUpdate().catch(console.error);
