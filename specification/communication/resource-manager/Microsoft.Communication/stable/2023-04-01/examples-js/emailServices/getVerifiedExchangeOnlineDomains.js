const { CommunicationServiceManagementClient } = require("@azure/arm-communication");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a list of domains that are fully verified in Exchange Online.
 *
 * @summary Get a list of domains that are fully verified in Exchange Online.
 * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/stable/2023-04-01/examples/emailServices/getVerifiedExchangeOnlineDomains.json
 */
async function getVerifiedExchangeOnlineDomains() {
  const subscriptionId =
    process.env["COMMUNICATION_SUBSCRIPTION_ID"] || "11112222-3333-4444-5555-666677778888";
  const credential = new DefaultAzureCredential();
  const client = new CommunicationServiceManagementClient(credential, subscriptionId);
  const result = await client.emailServices.listVerifiedExchangeOnlineDomains();
  console.log(result);
}
