const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get an alert that is associated a resource group or a resource in a resource group
 *
 * @summary Get an alert that is associated a resource group or a resource in a resource group
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2021-11-01/examples/Alerts/GetAlertResourceGroupLocation_example.json
 */
async function getSecurityAlertOnAResourceGroupFromASecurityDataLocation() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const ascLocation = "westeurope";
  const alertName = "2518298467986649999_4d25bfef-2d77-4a08-adc0-3e35715cc92a";
  const resourceGroupName = "myRg1";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.alerts.getResourceGroupLevel(
    ascLocation,
    alertName,
    resourceGroupName
  );
  console.log(result);
}

getSecurityAlertOnAResourceGroupFromASecurityDataLocation().catch(console.error);
