const { MariaDBManagementClient } = require("@azure/arm-mariadb");
const { DefaultAzureCredential } = require("@azure/identity");

async function updateAServerThreatDetectionPolicyWithMinimalParameters() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "securityalert-4799";
  const serverName = "securityalert-6440";
  const securityAlertPolicyName = "Default";
  const parameters = {
    emailAccountAdmins: true,
    state: "Disabled",
  };
  const credential = new DefaultAzureCredential();
  const client = new MariaDBManagementClient(credential, subscriptionId);
  const result = await client.serverSecurityAlertPolicies.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serverName,
    securityAlertPolicyName,
    parameters
  );
  console.log(result);
}

updateAServerThreatDetectionPolicyWithMinimalParameters().catch(console.error);
