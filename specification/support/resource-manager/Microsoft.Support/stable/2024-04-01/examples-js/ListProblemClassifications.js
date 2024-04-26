const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the problem classifications (categories) available for a specific Azure service. Always use the service and problem classifications obtained programmatically. This practice ensures that you always have the most recent set of service and problem classification Ids.
 *
 * @summary Lists all the problem classifications (categories) available for a specific Azure service. Always use the service and problem classifications obtained programmatically. This practice ensures that you always have the most recent set of service and problem classification Ids.
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/ListProblemClassifications.json
 */
async function getsListOfProblemClassificationsForAServiceForWhichASupportTicketCanBeCreated() {
  const serviceName = "service_guid";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential);
  const resArray = new Array();
  for await (let item of client.problemClassifications.list(serviceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
