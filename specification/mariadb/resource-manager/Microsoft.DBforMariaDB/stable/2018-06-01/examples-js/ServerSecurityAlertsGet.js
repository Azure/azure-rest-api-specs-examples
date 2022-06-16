const { MariaDBManagementClient } = require("@azure/arm-mariadb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a server's security alert policy.
 *
 * @summary Get a server's security alert policy.
 * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ServerSecurityAlertsGet.json
 */
async function getAServerThreatDetectionPolicy() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "securityalert-4799";
  const serverName = "securityalert-6440";
  const securityAlertPolicyName = "Default";
  const credential = new DefaultAzureCredential();
  const client = new MariaDBManagementClient(credential, subscriptionId);
  const result = await client.serverSecurityAlertPolicies.get(
    resourceGroupName,
    serverName,
    securityAlertPolicyName
  );
  console.log(result);
}

getAServerThreatDetectionPolicy().catch(console.error);
