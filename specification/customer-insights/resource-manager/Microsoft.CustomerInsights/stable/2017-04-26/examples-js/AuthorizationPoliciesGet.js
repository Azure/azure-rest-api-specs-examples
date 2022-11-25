const { CustomerInsightsManagementClient } = require("@azure/arm-customerinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an authorization policy in the hub.
 *
 * @summary Gets an authorization policy in the hub.
 * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/AuthorizationPoliciesGet.json
 */
async function authorizationPoliciesGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "TestHubRG";
  const hubName = "azSdkTestHub";
  const authorizationPolicyName = "testPolicy4222";
  const credential = new DefaultAzureCredential();
  const client = new CustomerInsightsManagementClient(credential, subscriptionId);
  const result = await client.authorizationPolicies.get(
    resourceGroupName,
    hubName,
    authorizationPolicyName
  );
  console.log(result);
}

authorizationPoliciesGet().catch(console.error);
