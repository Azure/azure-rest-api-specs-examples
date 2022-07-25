const { CommunicationServiceManagementClient } = require("@azure/arm-communication");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a list of domains that are fully verified in Exchange Online.
 *
 * @summary Get a list of domains that are fully verified in Exchange Online.
 * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2021-10-01-preview/examples/emailServices/getVerifiedExchangeOnlineDomains.json
 */
async function getVerifiedExchangeOnlineDomains() {
  const subscriptionId = "12345";
  const credential = new DefaultAzureCredential();
  const client = new CommunicationServiceManagementClient(credential, subscriptionId);
  const result = await client.emailServices.listVerifiedExchangeOnlineDomains();
  console.log(result);
}

getVerifiedExchangeOnlineDomains().catch(console.error);
