const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or update policy with specified rule set name within a resource group.
 *
 * @summary Creates or update policy with specified rule set name within a resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/WafPolicyCreateOrUpdate.json
 */
async function createsOrUpdatesAWafPolicyWithinAResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const policyName = "Policy1";
  const parameters = {
    customRules: [
      {
        name: "Rule1",
        action: "Block",
        matchConditions: [
          {
            matchValues: ["192.168.1.0/24", "10.0.0.0/24"],
            matchVariables: [{ selector: undefined, variableName: "RemoteAddr" }],
            operator: "IPMatch",
          },
        ],
        priority: 1,
        ruleType: "MatchRule",
      },
      {
        name: "Rule2",
        action: "Block",
        matchConditions: [
          {
            matchValues: ["192.168.1.0/24"],
            matchVariables: [{ selector: undefined, variableName: "RemoteAddr" }],
            operator: "IPMatch",
          },
          {
            matchValues: ["Windows"],
            matchVariables: [{ selector: "UserAgent", variableName: "RequestHeaders" }],
            operator: "Contains",
          },
        ],
        priority: 2,
        ruleType: "MatchRule",
      },
    ],
    location: "WestUs",
    managedRules: {
      exclusions: [
        {
          exclusionManagedRuleSets: [
            {
              ruleGroups: [
                {
                  ruleGroupName: "REQUEST-930-APPLICATION-ATTACK-LFI",
                  rules: [{ ruleId: "930120" }],
                },
                { ruleGroupName: "REQUEST-932-APPLICATION-ATTACK-RCE" },
              ],
              ruleSetType: "OWASP",
              ruleSetVersion: "3.2",
            },
          ],
          matchVariable: "RequestArgNames",
          selector: "hello",
          selectorMatchOperator: "StartsWith",
        },
        {
          exclusionManagedRuleSets: [
            { ruleGroups: [], ruleSetType: "OWASP", ruleSetVersion: "3.1" },
          ],
          matchVariable: "RequestArgNames",
          selector: "hello",
          selectorMatchOperator: "EndsWith",
        },
        {
          matchVariable: "RequestArgNames",
          selector: "test",
          selectorMatchOperator: "StartsWith",
        },
        {
          matchVariable: "RequestArgValues",
          selector: "test",
          selectorMatchOperator: "StartsWith",
        },
      ],
      managedRuleSets: [{ ruleSetType: "OWASP", ruleSetVersion: "3.2" }],
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.webApplicationFirewallPolicies.createOrUpdate(
    resourceGroupName,
    policyName,
    parameters
  );
  console.log(result);
}

createsOrUpdatesAWafPolicyWithinAResourceGroup().catch(console.error);
