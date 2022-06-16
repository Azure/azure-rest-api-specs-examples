const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all communications (attachments not included) for a support ticket. <br/></br> You can also filter support ticket communications by _CreatedDate_ or _CommunicationType_ using the $filter parameter. The only type of communication supported today is _Web_. Output will be a paged result with _nextLink_, using which you can retrieve the next set of Communication results. <br/><br/>Support ticket data is available for 18 months after ticket creation. If a ticket was created more than 18 months ago, a request for data might cause an error.
 *
 * @summary Lists all communications (attachments not included) for a support ticket. <br/></br> You can also filter support ticket communications by _CreatedDate_ or _CommunicationType_ using the $filter parameter. The only type of communication supported today is _Web_. Output will be a paged result with _nextLink_, using which you can retrieve the next set of Communication results. <br/><br/>Support ticket data is available for 18 months after ticket creation. If a ticket was created more than 18 months ago, a request for data might cause an error.
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/ListWebCommunicationsForSubscriptionSupportTicketCreatedOnOrAfter.json
 */
async function listWebCommunicationCreatedOnOrAfterASpecificDateForASubscriptionSupportTicket() {
  const subscriptionId = "subid";
  const supportTicketName = "testticket";
  const filter = "communicationType eq 'web' and createdDate ge 2020-03-10T22:08:51Z";
  const options = { filter };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.communications.list(supportTicketName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listWebCommunicationCreatedOnOrAfterASpecificDateForASubscriptionSupportTicket().catch(
  console.error
);
