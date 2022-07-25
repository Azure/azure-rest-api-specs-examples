const { AzureMLWebServicesManagementClient } = require("@azure/arm-webservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified web service.
 *
 * @summary Deletes the specified web service.
 * x-ms-original-file: specification/machinelearning/resource-manager/Microsoft.MachineLearning/stable/2017-01-01/examples/deleteWebService.json
 */
async function deleteWebService() {
  const subscriptionId = "subscription-id";
  const resourceGroupName = "OneResourceGroupName";
  const webServiceName = "TargetWebServiceName";
  const credential = new DefaultAzureCredential();
  const client = new AzureMLWebServicesManagementClient(credential, subscriptionId);
  const result = await client.webServices.beginRemoveAndWait(resourceGroupName, webServiceName);
  console.log(result);
}

deleteWebService().catch(console.error);
