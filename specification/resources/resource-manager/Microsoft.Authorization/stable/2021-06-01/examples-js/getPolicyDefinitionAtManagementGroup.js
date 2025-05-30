const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation retrieves the policy definition in the given management group with the given name.
 *
 * @summary This operation retrieves the policy definition in the given management group with the given name.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/getPolicyDefinitionAtManagementGroup.json
 */
async function retrieveAPolicyDefinitionAtManagementGroupLevel() {
  const subscriptionId =
    process.env["POLICY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const policyDefinitionName = "ResourceNaming";
  const managementGroupId = "MyManagementGroup";
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential, subscriptionId);
  const result = await client.policyDefinitions.getAtManagementGroup(
    policyDefinitionName,
    managementGroupId,
  );
  console.log(result);
}
