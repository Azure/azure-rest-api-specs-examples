const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets alerts for a resource scope.
 *
 * @summary Gets alerts for a resource scope.
 * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2022-08-01-preview/examples/GetAlerts.json
 */
async function getAlerts() {
  const scope = "subscriptions/afa2a084-766f-4003-8ae1-c4aeb893a99f";
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.alerts.listForScope(scope)) {
    resArray.push(item);
  }
  console.log(resArray);
}
