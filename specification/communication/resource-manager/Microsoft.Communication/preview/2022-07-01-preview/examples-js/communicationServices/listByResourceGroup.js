const { CommunicationServiceManagementClient } = require("@azure/arm-communication");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Handles requests to list all resources in a resource group.
 *
 * @summary Handles requests to list all resources in a resource group.
 * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2022-07-01-preview/examples/communicationServices/listByResourceGroup.json
 */
async function listByResourceGroup() {
  const subscriptionId = "12345";
  const resourceGroupName = "MyResourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new CommunicationServiceManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.communicationServices.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listByResourceGroup().catch(console.error);
