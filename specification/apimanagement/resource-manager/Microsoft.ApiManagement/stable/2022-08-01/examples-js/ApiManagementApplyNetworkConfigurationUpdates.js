const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the Microsoft.ApiManagement resource running in the Virtual network to pick the updated DNS changes.
 *
 * @summary Updates the Microsoft.ApiManagement resource running in the Virtual network to pick the updated DNS changes.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementApplyNetworkConfigurationUpdates.json
 */
async function apiManagementApplyNetworkConfigurationUpdates() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const parameters = {
    location: "west us",
  };
  const options = {
    parameters,
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiManagementService.beginApplyNetworkConfigurationUpdatesAndWait(
    resourceGroupName,
    serviceName,
    options
  );
  console.log(result);
}
