const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to This operation retrieves the list of all policy assignments associated with the given resource group in the given subscription that match the optional given $filter. Valid values for $filter are: 'atScope()', 'atExactScope()' or 'policyDefinitionId eq '{value}''. If $filter is not provided, the unfiltered list includes all policy assignments associated with the resource group, including those that apply directly or apply from containing scopes, as well as any applied to resources contained within the resource group. If $filter=atScope() is provided, the returned list includes all policy assignments that apply to the resource group, which is everything in the unfiltered list except those applied to resources contained within the resource group. If $filter=atExactScope() is provided, the returned list only includes all policy assignments that at the resource group. If $filter=policyDefinitionId eq '{value}' is provided, the returned list includes all policy assignments of the policy definition whose id is {value} that apply to the resource group.
 *
 * @summary This operation retrieves the list of all policy assignments associated with the given resource group in the given subscription that match the optional given $filter. Valid values for $filter are: 'atScope()', 'atExactScope()' or 'policyDefinitionId eq '{value}''. If $filter is not provided, the unfiltered list includes all policy assignments associated with the resource group, including those that apply directly or apply from containing scopes, as well as any applied to resources contained within the resource group. If $filter=atScope() is provided, the returned list includes all policy assignments that apply to the resource group, which is everything in the unfiltered list except those applied to resources contained within the resource group. If $filter=atExactScope() is provided, the returned list only includes all policy assignments that at the resource group. If $filter=policyDefinitionId eq '{value}' is provided, the returned list includes all policy assignments of the policy definition whose id is {value} that apply to the resource group.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/listPolicyAssignmentsForResourceGroup.json
 */
async function listPolicyAssignmentsThatApplyToAResourceGroup() {
  const subscriptionId =
    process.env["POLICY_SUBSCRIPTION_ID"] || "ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
  const resourceGroupName = process.env["POLICY_RESOURCE_GROUP"] || "TestResourceGroup";
  const filter = "atScope()";
  const expand = "LatestDefinitionVersion, EffectiveDefinitionVersion";
  const options = {
    filter,
    expand,
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.policyAssignments.listForResourceGroup(
    resourceGroupName,
    options,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
