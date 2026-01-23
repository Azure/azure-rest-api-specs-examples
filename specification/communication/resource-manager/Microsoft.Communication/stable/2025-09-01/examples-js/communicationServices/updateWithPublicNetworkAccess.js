const { CommunicationServiceManagementClient } = require("@azure/arm-communication");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Operation to update an existing CommunicationService.
 *
 * @summary Operation to update an existing CommunicationService.
 * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/communicationServices/updateWithPublicNetworkAccess.json
 */
async function updateResourceToAddPublicNetworkAccess() {
  const subscriptionId =
    process.env["COMMUNICATION_SUBSCRIPTION_ID"] || "11112222-3333-4444-5555-666677778888";
  const resourceGroupName = process.env["COMMUNICATION_RESOURCE_GROUP"] || "MyResourceGroup";
  const communicationServiceName = "MyCommunicationResource";
  const parameters = {
    publicNetworkAccess: "Enabled",
    tags: { newTag: "newVal" },
  };
  const credential = new DefaultAzureCredential();
  const client = new CommunicationServiceManagementClient(credential, subscriptionId);
  const result = await client.communicationServices.update(
    resourceGroupName,
    communicationServiceName,
    parameters,
  );
  console.log(result);
}
