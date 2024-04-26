const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the support tickets. <br/><br/>You can also filter the support tickets by <i>Status</i>, <i>CreatedDate</i>, , <i>ServiceId</i>, and <i>ProblemClassificationId</i> using the $filter parameter. Output will be a paged result with <i>nextLink</i>, using which you can retrieve the next set of support tickets. <br/><br/>Support ticket data is available for 18 months after ticket creation. If a ticket was created more than 18 months ago, a request for data might cause an error.
 *
 * @summary Lists all the support tickets. <br/><br/>You can also filter the support tickets by <i>Status</i>, <i>CreatedDate</i>, , <i>ServiceId</i>, and <i>ProblemClassificationId</i> using the $filter parameter. Output will be a paged result with <i>nextLink</i>, using which you can retrieve the next set of support tickets. <br/><br/>Support ticket data is available for 18 months after ticket creation. If a ticket was created more than 18 months ago, a request for data might cause an error.
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/ListSupportTicketsCreatedOnOrAfterAndInUpdatingState.json
 */
async function listSupportTicketsCreatedOnOrAfterACertainDateAndInUpdatingState() {
  const filter = "createdDate ge 2020-03-10T22:08:51Z and status eq 'Updating'";
  const options = { filter };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential);
  const resArray = new Array();
  for await (let item of client.supportTicketsNoSubscription.list(options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
