const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This operation retrieves the policy set definition in the given management group with the given name.
 *
 * @summary This operation retrieves the policy set definition in the given management group with the given name.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/getPolicySetDefinitionAtManagementGroup.json
 */
async function retrieveAPolicySetDefinitionAtManagementGroupLevel() {
  const subscriptionId =
    process.env["POLICY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const policySetDefinitionName = "CostManagement";
  const managementGroupId = "MyManagementGroup";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential, subscriptionId);
  const result = await client.policySetDefinitions.getAtManagementGroup(
    policySetDefinitionName,
    managementGroupId
  );
  console.log(result);
}
