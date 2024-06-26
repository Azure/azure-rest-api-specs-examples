const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check if the Application Configuration Service settings are valid.
 *
 * @summary Check if the Application Configuration Service settings are valid.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/ConfigurationServices_Validate.json
 */
async function configurationServicesValidate() {
  const subscriptionId =
    process.env["APPPLATFORM_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPPLATFORM_RESOURCE_GROUP"] || "myResourceGroup";
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
    settings,
  );
  console.log(result);
}
