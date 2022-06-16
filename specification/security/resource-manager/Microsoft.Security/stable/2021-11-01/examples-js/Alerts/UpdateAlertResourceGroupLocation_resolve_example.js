const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the alert's state
 *
 * @summary Update the alert's state
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2021-11-01/examples/Alerts/UpdateAlertResourceGroupLocation_resolve_example.json
 */
async function updateSecurityAlertStateOnAResourceGroupFromASecurityDataLocation() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const ascLocation = "westeurope";
  const alertName = "2518765996949954086_2325cf9e-42a2-4f72-ae7f-9b863cba2d22";
  const resourceGroupName = "myRg2";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.alerts.updateResourceGroupLevelStateToResolve(
    ascLocation,
    alertName,
    resourceGroupName
  );
  console.log(result);
}

updateSecurityAlertStateOnAResourceGroupFromASecurityDataLocation().catch(console.error);
