const { AppConfigurationManagementClient } = require("@azure/arm-appconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets the properties of the specified key-value. NOTE: This operation is intended for use in ARM Template deployments. For all other scenarios involving App Configuration key-values the data plane API should be used instead.
 *
 * @summary Gets the properties of the specified key-value. NOTE: This operation is intended for use in ARM Template deployments. For all other scenarios involving App Configuration key-values the data plane API should be used instead.
 * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2024-06-01/examples/ConfigurationStoresGetKeyValue.json
 */
async function keyValuesGet() {
  const subscriptionId =
    process.env["APPCONFIGURATION_SUBSCRIPTION_ID"] || "c80fb759-c965-4c6a-9110-9b2b2d038882";
  const resourceGroupName = process.env["APPCONFIGURATION_RESOURCE_GROUP"] || "myResourceGroup";
  const configStoreName = "contoso";
  const keyValueName = "myKey$myLabel";
  const credential = new DefaultAzureCredential();
  const client = new AppConfigurationManagementClient(credential, subscriptionId);
  const result = await client.keyValues.get(resourceGroupName, configStoreName, keyValueName);
  console.log(result);
}
