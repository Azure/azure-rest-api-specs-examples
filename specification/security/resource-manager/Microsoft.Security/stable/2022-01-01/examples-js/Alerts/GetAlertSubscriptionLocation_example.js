const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get an alert that is associated with a subscription
 *
 * @summary Get an alert that is associated with a subscription
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2022-01-01/examples/Alerts/GetAlertSubscriptionLocation_example.json
 */
async function getSecurityAlertOnASubscriptionFromASecurityDataLocation() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const ascLocation = "westeurope";
  const alertName = "2518770965529163669_F144EE95-A3E5-42DA-A279-967D115809AA";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.alerts.getSubscriptionLevel(ascLocation, alertName);
  console.log(result);
}

getSecurityAlertOnASubscriptionFromASecurityDataLocation().catch(console.error);
