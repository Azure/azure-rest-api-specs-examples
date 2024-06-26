const { CustomerInsightsManagementClient } = require("@azure/arm-customerinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerates the primary policy key of the specified authorization policy.
 *
 * @summary Regenerates the primary policy key of the specified authorization policy.
 * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/AuthorizationPoliciesRegeneratePrimaryKey.json
 */
async function authorizationPoliciesRegeneratePrimaryKey() {
  const subscriptionId = "subid";
  const resourceGroupName = "TestHubRG";
  const hubName = "azSdkTestHub";
  const authorizationPolicyName = "testPolicy4222";
  const credential = new DefaultAzureCredential();
  const client = new CustomerInsightsManagementClient(credential, subscriptionId);
  const result = await client.authorizationPolicies.regeneratePrimaryKey(
    resourceGroupName,
    hubName,
    authorizationPolicyName
  );
  console.log(result);
}

authorizationPoliciesRegeneratePrimaryKey().catch(console.error);
