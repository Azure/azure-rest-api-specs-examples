const { CommunicationServiceManagementClient } = require("@azure/arm-communication");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks that the CommunicationService name is valid and is not already in use.
 *
 * @summary Checks that the CommunicationService name is valid and is not already in use.
 * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/stable/2023-03-31/examples/communicationServices/checkNameAvailabilityAvailable.json
 */
async function checkNameAvailabilityAvailable() {
  const subscriptionId =
    process.env["COMMUNICATION_SUBSCRIPTION_ID"] || "11112222-3333-4444-5555-666677778888";
  const nameAvailabilityParameters = {
    name: "MyCommunicationService",
    type: "Microsoft.Communication/CommunicationServices",
  };
  const credential = new DefaultAzureCredential();
  const client = new CommunicationServiceManagementClient(credential, subscriptionId);
  const result = await client.communicationServices.checkNameAvailability(
    nameAvailabilityParameters
  );
  console.log(result);
}
