const { CommunicationServiceManagementClient } = require("@azure/arm-communication");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to operation to update an existing CommunicationService.
 *
 * @summary operation to update an existing CommunicationService.
 * x-ms-original-file: 2026-03-18/communicationServices/updateWithDisableLocalAuth.json
 */
async function updateResourceToAddDisableLocalAuth() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "11112222-3333-4444-5555-666677778888";
  const client = new CommunicationServiceManagementClient(credential, subscriptionId);
  const result = await client.communicationServices.update(
    "MyResourceGroup",
    "MyCommunicationResource",
    { disableLocalAuth: true, tags: { newTag: "newVal" } },
  );
  console.log(result);
}
