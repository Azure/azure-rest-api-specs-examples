const { CommunicationServiceManagementClient } = require("@azure/arm-communication");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the CommunicationService and its properties.
 *
 * @summary Get the CommunicationService and its properties.
 * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2023-04-01-preview/examples/communicationServices/get.json
 */
async function getResource() {
  const subscriptionId =
    process.env["COMMUNICATION_SUBSCRIPTION_ID"] || "11112222-3333-4444-5555-666677778888";
  const resourceGroupName = process.env["COMMUNICATION_RESOURCE_GROUP"] || "MyResourceGroup";
  const communicationServiceName = "MyCommunicationResource";
  const credential = new DefaultAzureCredential();
  const client = new CommunicationServiceManagementClient(credential, subscriptionId);
  const result = await client.communicationServices.get(
    resourceGroupName,
    communicationServiceName
  );
  console.log(result);
}
