const { AutomanageClient } = require("@azure/arm-automanage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve a list of configuration profile within a given resource group
 *
 * @summary Retrieve a list of configuration profile within a given resource group
 * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listConfigurationProfilesByResourceGroup.json
 */
async function listConfigurationProfilesByResourceGroup() {
  const subscriptionId = "mySubscriptionId";
  const resourceGroupName = "myResourceGroupName";
  const credential = new DefaultAzureCredential();
  const client = new AutomanageClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.configurationProfiles.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listConfigurationProfilesByResourceGroup().catch(console.error);
