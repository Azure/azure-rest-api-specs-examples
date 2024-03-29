const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks what restrictions Azure Policy will place on resources within a management group.
 *
 * @summary Checks what restrictions Azure Policy will place on resources within a management group.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-03-01/examples/PolicyRestrictions_CheckAtManagementGroupScope.json
 */
async function checkPolicyRestrictionsAtManagementGroupScope() {
  const subscriptionId =
    process.env["POLICYINSIGHTS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const managementGroupId = "financeMg";
  const parameters = {
    pendingFields: [{ field: "type" }],
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const result = await client.policyRestrictions.checkAtManagementGroupScope(
    managementGroupId,
    parameters
  );
  console.log(result);
}
