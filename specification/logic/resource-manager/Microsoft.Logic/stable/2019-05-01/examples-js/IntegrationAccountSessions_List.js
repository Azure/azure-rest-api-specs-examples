const { LogicManagementClient } = require("@azure/arm-logic");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of integration account sessions.
 *
 * @summary Gets a list of integration account sessions.
 * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountSessions_List.json
 */
async function getAListOfIntegrationAccountSessions() {
  const subscriptionId =
    process.env["LOGIC_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["LOGIC_RESOURCE_GROUP"] || "testrg123";
  const integrationAccountName = "testia123";
  const credential = new DefaultAzureCredential();
  const client = new LogicManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.integrationAccountSessions.list(
    resourceGroupName,
    integrationAccountName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
