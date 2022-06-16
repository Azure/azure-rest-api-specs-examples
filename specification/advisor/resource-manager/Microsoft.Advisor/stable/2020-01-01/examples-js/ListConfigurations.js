const { AdvisorManagementClient } = require("@azure/arm-advisor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve Azure Advisor configurations.
 *
 * @summary Retrieve Azure Advisor configurations.
 * x-ms-original-file: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/ListConfigurations.json
 */
async function getConfigurations() {
  const subscriptionId = "subscriptionId";
  const resourceGroup = "resourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new AdvisorManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.configurations.listByResourceGroup(resourceGroup)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getConfigurations().catch(console.error);
