const { CommunicationServiceManagementClient } = require("@azure/arm-communication");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks that the CommunicationService name is valid and is not already in use.
 *
 * @summary Checks that the CommunicationService name is valid and is not already in use.
 * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/stable/2020-08-20/examples/checkNameAvailabilityUnavailable.json
 */
async function checkNameAvailabilityUnavailable() {
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

checkNameAvailabilityUnavailable().catch(console.error);
