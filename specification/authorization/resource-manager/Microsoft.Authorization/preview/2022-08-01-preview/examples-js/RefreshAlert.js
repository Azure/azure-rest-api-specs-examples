const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Refresh an alert.
 *
 * @summary Refresh an alert.
 * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2022-08-01-preview/examples/RefreshAlert.json
 */
async function refreshAlert() {
  const scope = "subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f";
  const alertId = "AzureRolesAssignedOutsidePimAlert";
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential);
  const result = await client.alerts.beginRefreshAndWait(scope, alertId);
  console.log(result);
}
