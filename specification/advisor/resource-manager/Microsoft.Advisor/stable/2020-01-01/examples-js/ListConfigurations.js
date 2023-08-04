const { AdvisorManagementClient } = require("@azure/arm-advisor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve Azure Advisor configurations and also retrieve configurations of contained resource groups.
 *
 * @summary Retrieve Azure Advisor configurations and also retrieve configurations of contained resource groups.
 * x-ms-original-file: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/ListConfigurations.json
 */
async function getConfigurations() {
  const subscriptionId = process.env["ADVISOR_SUBSCRIPTION_ID"] || "subscriptionId";
  const credential = new DefaultAzureCredential();
  const client = new AdvisorManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.configurations.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
