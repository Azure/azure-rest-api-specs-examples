const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the Shared Access Authorization Token for the User.
 *
 * @summary Gets the Shared Access Authorization Token for the User.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUserToken.json
 */
async function apiManagementUserToken() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const userId = "userId1718";
  const parameters = {
    expiry: new Date("2019-04-21T00:44:24.2845269Z"),
    keyType: "primary",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.user.getSharedAccessToken(
    resourceGroupName,
    serviceName,
    userId,
    parameters
  );
  console.log(result);
}
