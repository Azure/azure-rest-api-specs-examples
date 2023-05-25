const { PolicyClient } = require("@azure/arm-policy");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This operation retrieves the list of all policy exemptions associated with the given resource group in the given subscription that match the optional given $filter. Valid values for $filter are: 'atScope()', 'atExactScope()', 'excludeExpired()' or 'policyAssignmentId eq '{value}''. If $filter is not provided, the unfiltered list includes all policy exemptions associated with the resource group, including those that apply directly or apply from containing scopes, as well as any applied to resources contained within the resource group.
 *
 * @summary This operation retrieves the list of all policy exemptions associated with the given resource group in the given subscription that match the optional given $filter. Valid values for $filter are: 'atScope()', 'atExactScope()', 'excludeExpired()' or 'policyAssignmentId eq '{value}''. If $filter is not provided, the unfiltered list includes all policy exemptions associated with the resource group, including those that apply directly or apply from containing scopes, as well as any applied to resources contained within the resource group.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/preview/2020-07-01-preview/examples/listPolicyExemptionsForResourceGroup.json
 */
async function listPolicyExemptionsThatApplyToAResourceGroup() {
  const subscriptionId =
    process.env["POLICY_SUBSCRIPTION_ID"] || "ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
  const resourceGroupName = process.env["POLICY_RESOURCE_GROUP"] || "TestResourceGroup";
  const filter = "atScope()";
  const options = {
    filter,
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.policyExemptions.listForResourceGroup(resourceGroupName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
