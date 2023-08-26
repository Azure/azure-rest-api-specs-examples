const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the collection of subscriptions of the specified user.
 *
 * @summary Lists the collection of subscriptions of the specified user.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListUserSubscriptions.json
 */
async function apiManagementListUserSubscriptions() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const userId = "57681833a40f7eb6c49f6acf";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.userSubscription.list(resourceGroupName, serviceName, userId)) {
    resArray.push(item);
  }
  console.log(resArray);
}
