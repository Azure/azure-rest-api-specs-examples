const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the entity state (Etag) version of the apimanagement subscription specified by its identifier.
 *
 * @summary Gets the entity state (Etag) version of the apimanagement subscription specified by its identifier.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadSubscription.json
 */
async function apiManagementHeadSubscription() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const sid = "5931a769d8d14f0ad8ce13b8";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.subscription.getEntityTag(resourceGroupName, serviceName, sid);
  console.log(result);
}
