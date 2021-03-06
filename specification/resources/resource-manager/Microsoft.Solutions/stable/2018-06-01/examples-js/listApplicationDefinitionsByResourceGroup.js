const { ApplicationClient } = require("@azure/arm-managedapplications");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the managed application definitions in a resource group.
 *
 * @summary Lists the managed application definitions in a resource group.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/listApplicationDefinitionsByResourceGroup.json
 */
async function listManagedApplicationDefinitions() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.applicationDefinitions.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listManagedApplicationDefinitions().catch(console.error);
