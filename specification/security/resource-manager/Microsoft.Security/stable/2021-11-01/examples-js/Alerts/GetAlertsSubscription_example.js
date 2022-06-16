const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all the alerts that are associated with the subscription
 *
 * @summary List all the alerts that are associated with the subscription
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2021-11-01/examples/Alerts/GetAlertsSubscription_example.json
 */
async function getSecurityAlertsOnASubscription() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.alerts.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

getSecurityAlertsOnASubscription().catch(console.error);
