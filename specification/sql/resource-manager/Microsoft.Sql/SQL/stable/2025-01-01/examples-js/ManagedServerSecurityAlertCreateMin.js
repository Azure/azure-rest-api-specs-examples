const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a threat detection policy.
 *
 * @summary creates or updates a threat detection policy.
 * x-ms-original-file: 2025-01-01/ManagedServerSecurityAlertCreateMin.json
 */
async function updateAManagedServerThreatDetectionPolicyWithMinimalParameters() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.managedServerSecurityAlertPolicies.createOrUpdate(
    "securityalert-4799",
    "securityalert-6440",
    "Default",
    { state: "Enabled" },
  );
  console.log(result);
}
