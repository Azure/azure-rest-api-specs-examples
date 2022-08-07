const { AutomanageClient } = require("@azure/arm-automanage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve a list of configuration profile version for a configuration profile
 *
 * @summary Retrieve a list of configuration profile version for a configuration profile
 * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listConfigurationProfileVersions.json
 */
async function listConfigurationProfileVersionsByConfigurationProfile() {
  const subscriptionId = "mySubscriptionId";
  const configurationProfileName = "customConfigurationProfile";
  const resourceGroupName = "myResourceGroupName";
  const credential = new DefaultAzureCredential();
  const client = new AutomanageClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.configurationProfilesVersions.listChildResources(
    configurationProfileName,
    resourceGroupName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listConfigurationProfileVersionsByConfigurationProfile().catch(console.error);
