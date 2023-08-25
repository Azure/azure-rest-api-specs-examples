const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the list of Microsoft.ApiManagement SKUs available for your Subscription.
 *
 * @summary Gets the list of Microsoft.ApiManagement SKUs available for your Subscription.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListSku.json
 */
async function listsAllAvailableResourceSkUs() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.apiManagementSkus.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
