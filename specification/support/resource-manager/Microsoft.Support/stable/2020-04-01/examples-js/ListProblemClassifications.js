const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

async function getsListOfProblemClassificationsForAServiceForWhichASupportTicketCanBeCreated() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const serviceName = "service_guid";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.problemClassifications.list(serviceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getsListOfProblemClassificationsForAServiceForWhichASupportTicketCanBeCreated().catch(
  console.error
);
