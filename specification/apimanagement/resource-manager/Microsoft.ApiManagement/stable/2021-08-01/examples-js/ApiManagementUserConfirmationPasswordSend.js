const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Sends confirmation
 *
 * @summary Sends confirmation
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUserConfirmationPasswordSend.json
 */
async function apiManagementUserConfirmationPasswordSend() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const userId = "57127d485157a511ace86ae7";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.userConfirmationPassword.send(resourceGroupName, serviceName, userId);
  console.log(result);
}

apiManagementUserConfirmationPasswordSend().catch(console.error);
