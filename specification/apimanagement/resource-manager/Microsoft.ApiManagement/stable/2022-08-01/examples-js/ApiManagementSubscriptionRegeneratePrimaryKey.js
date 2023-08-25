const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerates primary key of existing subscription of the API Management service instance.
 *
 * @summary Regenerates primary key of existing subscription of the API Management service instance.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementSubscriptionRegeneratePrimaryKey.json
 */
async function apiManagementSubscriptionRegeneratePrimaryKey() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const sid = "testsub";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.subscription.regeneratePrimaryKey(
    resourceGroupName,
    serviceName,
    sid
  );
  console.log(result);
}
