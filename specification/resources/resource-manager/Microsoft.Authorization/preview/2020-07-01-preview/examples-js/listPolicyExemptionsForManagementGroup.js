const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This operation retrieves the list of all policy exemptions applicable to the management group that match the given $filter. Valid values for $filter are: 'atScope()', 'atExactScope()', 'excludeExpired()' or 'policyAssignmentId eq '{value}''. If $filter=atScope() is provided, the returned list includes all policy exemptions that are assigned to the management group or the management group's ancestors.
 *
 * @summary This operation retrieves the list of all policy exemptions applicable to the management group that match the given $filter. Valid values for $filter are: 'atScope()', 'atExactScope()', 'excludeExpired()' or 'policyAssignmentId eq '{value}''. If $filter=atScope() is provided, the returned list includes all policy exemptions that are assigned to the management group or the management group's ancestors.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/preview/2020-07-01-preview/examples/listPolicyExemptionsForManagementGroup.json
 */
async function listPolicyExemptionsThatApplyToAManagementGroup() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const managementGroupId = "DevOrg";
  const filter = "atScope()";
  const options = {
    filter,
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.policyExemptions.listForManagementGroup(
    managementGroupId,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listPolicyExemptionsThatApplyToAManagementGroup().catch(console.error);
