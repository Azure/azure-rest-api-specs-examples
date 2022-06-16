const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check if the Application Configuration Service settings are valid.
 *
 * @summary Check if the Application Configuration Service settings are valid.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/ConfigurationServices_Validate.json
 */
async function configurationServicesValidate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const configurationServiceName = "default";
  const settings = {
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
  };
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.configurationServices.beginValidateAndWait(
    resourceGroupName,
    serviceName,
    configurationServiceName,
    settings
  );
  console.log(result);
}

configurationServicesValidate().catch(console.error);
