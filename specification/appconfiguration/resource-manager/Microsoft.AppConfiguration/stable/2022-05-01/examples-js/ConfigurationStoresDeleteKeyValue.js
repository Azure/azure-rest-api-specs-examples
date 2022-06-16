const { AppConfigurationManagementClient } = require("@azure/arm-appconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a key-value.
 *
 * @summary Deletes a key-value.
 * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2022-05-01/examples/ConfigurationStoresDeleteKeyValue.json
 */
async function keyValuesDelete() {
  const subscriptionId = "c80fb759-c965-4c6a-9110-9b2b2d038882";
  const resourceGroupName = "myResourceGroup";
  const configStoreName = "contoso";
  const keyValueName = "myKey$myLabel";
  const credential = new DefaultAzureCredential();
  const client = new AppConfigurationManagementClient(credential, subscriptionId);
  const result = await client.keyValues.beginDeleteAndWait(
    resourceGroupName,
    configStoreName,
    keyValueName
  );
  console.log(result);
}

keyValuesDelete().catch(console.error);
