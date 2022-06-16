const { AppConfigurationManagementClient } = require("@azure/arm-appconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a key-value.
 *
 * @summary Creates a key-value.
 * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2022-05-01/examples/ConfigurationStoresCreateKeyValue.json
 */
async function keyValuesCreateOrUpdate() {
  const subscriptionId = "c80fb759-c965-4c6a-9110-9b2b2d038882";
  const resourceGroupName = "myResourceGroup";
  const configStoreName = "contoso";
  const keyValueName = "myKey$myLabel";
  const keyValueParameters = {
    tags: { tag1: "tagValue1", tag2: "tagValue2" },
    value: "myValue",
  };
  const options = { keyValueParameters };
  const credential = new DefaultAzureCredential();
  const client = new AppConfigurationManagementClient(credential, subscriptionId);
  const result = await client.keyValues.createOrUpdate(
    resourceGroupName,
    configStoreName,
    keyValueName,
    options
  );
  console.log(result);
}

keyValuesCreateOrUpdate().catch(console.error);
