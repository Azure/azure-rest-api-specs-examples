const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the support tickets for an Azure subscription. You can also filter the support tickets by _Status_ or _CreatedDate_ using the $filter parameter. Output will be a paged result with _nextLink_, using which you can retrieve the next set of support tickets. <br/><br/>Support ticket data is available for 18 months after ticket creation. If a ticket was created more than 18 months ago, a request for data might cause an error.
 *
 * @summary Lists all the support tickets for an Azure subscription. You can also filter the support tickets by _Status_ or _CreatedDate_ using the $filter parameter. Output will be a paged result with _nextLink_, using which you can retrieve the next set of support tickets. <br/><br/>Support ticket data is available for 18 months after ticket creation. If a ticket was created more than 18 months ago, a request for data might cause an error.
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/ListSupportTicketsCreatedOnOrAfterAndInOpenStateBySubscription.json
 */
async function listSupportTicketsCreatedOnOrAfterACertainDateAndInOpenStateForASubscription() {
  const subscriptionId = "subid";
  const filter = "createdDate ge 2020-03-10T22:08:51Z and status eq 'Open'";
  const options = { filter };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.supportTickets.list(options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listSupportTicketsCreatedOnOrAfterACertainDateAndInOpenStateForASubscription().catch(console.error);
