const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the subscription of specified user to the specified product.
 *
 * @summary Creates or updates the subscription of specified user to the specified product.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateSubscription.json
 */
async function apiManagementCreateSubscription() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const sid = "testsub";
  const parameters = {
    displayName: "testsub",
    ownerId:
      "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/57127d485157a511ace86ae7",
    scope:
      "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/products/5600b59475ff190048060002",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.subscription.createOrUpdate(
    resourceGroupName,
    serviceName,
    sid,
    parameters
  );
  console.log(result);
}
