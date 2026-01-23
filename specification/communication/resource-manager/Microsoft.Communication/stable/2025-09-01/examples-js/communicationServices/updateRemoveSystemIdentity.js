const { CommunicationServiceManagementClient } = require("@azure/arm-communication");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Operation to update an existing CommunicationService.
 *
 * @summary Operation to update an existing CommunicationService.
 * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/communicationServices/updateRemoveSystemIdentity.json
 */
async function updateResourceToRemoveIdentity() {
  const subscriptionId =
    process.env["COMMUNICATION_SUBSCRIPTION_ID"] || "11112222-3333-4444-5555-666677778888";
  const resourceGroupName = process.env["COMMUNICATION_RESOURCE_GROUP"] || "MyResourceGroup";
  const communicationServiceName = "MyCommunicationResource";
  const parameters = {
    identity: { type: "None" },
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
