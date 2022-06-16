const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing security policy within a profile.
 *
 * @summary Updates an existing security policy within a profile.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/SecurityPolicies_Patch.json
 */
async function securityPoliciesPatch() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const securityPolicyName = "securityPolicy1";
  const securityPolicyUpdateProperties = {
    parameters: {
      type: "WebApplicationFirewall",
      associations: [
        {
          domains: [
            {
              id: "/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/afddomains/testdomain1",
            },
            {
              id: "/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/afddomains/testdomain2",
            },
          ],
          patternsToMatch: ["/*"],
        },
      ],
      wafPolicy: {
        id: "/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Network/frontdoorwebapplicationfirewallpolicies/wafTest",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.securityPolicies.beginPatchAndWait(
    resourceGroupName,
    profileName,
    securityPolicyName,
    securityPolicyUpdateProperties
  );
  console.log(result);
}

securityPoliciesPatch().catch(console.error);
