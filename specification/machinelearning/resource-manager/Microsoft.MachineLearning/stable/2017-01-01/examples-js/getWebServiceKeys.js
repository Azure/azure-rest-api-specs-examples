const { AzureMLWebServicesManagementClient } = require("@azure/arm-webservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the access keys for the specified web service.
 *
 * @summary Gets the access keys for the specified web service.
 * x-ms-original-file: specification/machinelearning/resource-manager/Microsoft.MachineLearning/stable/2017-01-01/examples/getWebServiceKeys.json
 */
async function getWebServiceKeys() {
  const subscriptionId = "subscription-id";
  const resourceGroupName = "OneResourceGroupName";
  const webServiceName = "TargetWebServiceName";
  const credential = new DefaultAzureCredential();
  const client = new AzureMLWebServicesManagementClient(credential, subscriptionId);
  const result = await client.webServices.listKeys(resourceGroupName, webServiceName);
  console.log(result);
}

getWebServiceKeys().catch(console.error);
