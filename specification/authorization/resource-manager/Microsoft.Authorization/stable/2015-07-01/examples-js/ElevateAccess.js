const { AuthorizationManagementClient } = require("@azure/arm-authorization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Elevates access for a Global Administrator.
 *
 * @summary Elevates access for a Global Administrator.
 * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2015-07-01/examples/ElevateAccess.json
 */
async function elevateAccessGlobalAdministrator() {
  const credential = new DefaultAzureCredential();
  const client = new AuthorizationManagementClient(credential);
  const result = await client.globalAdministrator.elevateAccess();
  console.log(result);
}
