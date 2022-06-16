const { CommunicationServiceManagementClient } = require("@azure/arm-communication");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks that the CommunicationService name is valid and is not already in use.
 *
 * @summary Checks that the CommunicationService name is valid and is not already in use.
 * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/stable/2020-08-20/examples/checkNameAvailabilityAvailable.json
 */
async function checkNameAvailabilityAvailable() {
  const subscriptionId = "12345";
  const nameAvailabilityParameters = {
    name: "MyCommunicationService",
    type: "Microsoft.Communication/CommunicationServices",
  };
  const options = {
    nameAvailabilityParameters,
  };
  const credential = new DefaultAzureCredential();
  const client = new CommunicationServiceManagementClient(credential, subscriptionId);
  const result = await client.communicationService.checkNameAvailability(options);
  console.log(result);
}

checkNameAvailabilityAvailable().catch(console.error);
