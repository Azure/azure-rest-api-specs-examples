const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update policy with specified rule set name within a resource group.
 *
 * @summary Create or update policy with specified rule set name within a resource group.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/WafPolicyCreateOrUpdate.json
 */
async function createsSpecificPolicy() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const policyName = "MicrosoftCdnWafPolicy";
  const cdnWebApplicationFirewallPolicy = {
    customRules: {
      rules: [
        {
          name: "CustomRule1",
          action: "Block",
          enabledState: "Enabled",
          matchConditions: [
            {
              matchValue: ["CH"],
              matchVariable: "RemoteAddr",
              negateCondition: false,
              operator: "GeoMatch",
              selector: undefined,
              transforms: [],
            },
            {
              matchValue: ["windows"],
              matchVariable: "RequestHeader",
              negateCondition: false,
              operator: "Contains",
              selector: "UserAgent",
              transforms: [],
            },
            {
              matchValue: ["<?php", "?>"],
              matchVariable: "QueryString",
              negateCondition: false,
              operator: "Contains",
              selector: "search",
              transforms: ["UrlDecode", "Lowercase"],
            },
          ],
          priority: 2,
        },
      ],
    },
    location: "WestUs",
    managedRules: {
      managedRuleSets: [
        {
          ruleGroupOverrides: [
            {
              ruleGroupName: "Group1",
              rules: [
                {
                  action: "Redirect",
                  enabledState: "Enabled",
                  ruleId: "GROUP1-0001",
                },
                { enabledState: "Disabled", ruleId: "GROUP1-0002" },
              ],
            },
          ],
          ruleSetType: "DefaultRuleSet",
          ruleSetVersion: "preview-1.0",
        },
      ],
    },
    policySettings: {
      defaultCustomBlockResponseBody:
        "PGh0bWw+CjxoZWFkZXI+PHRpdGxlPkhlbGxvPC90aXRsZT48L2hlYWRlcj4KPGJvZHk+CkhlbGxvIHdvcmxkCjwvYm9keT4KPC9odG1sPg==",
      defaultCustomBlockResponseStatusCode: 200,
      defaultRedirectUrl: "http://www.bing.com",
    },
    rateLimitRules: {
      rules: [
        {
          name: "RateLimitRule1",
          action: "Block",
          enabledState: "Enabled",
          matchConditions: [
            {
              matchValue: ["192.168.1.0/24", "10.0.0.0/24"],
              matchVariable: "RemoteAddr",
              negateCondition: false,
              operator: "IPMatch",
              selector: undefined,
              transforms: [],
            },
          ],
          priority: 1,
          rateLimitDurationInMinutes: 0,
          rateLimitThreshold: 1000,
        },
      ],
    },
    sku: { name: "Standard_Microsoft" },
  };
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.policies.beginCreateOrUpdateAndWait(
    resourceGroupName,
    policyName,
    cdnWebApplicationFirewallPolicy
  );
  console.log(result);
}

createsSpecificPolicy().catch(console.error);
