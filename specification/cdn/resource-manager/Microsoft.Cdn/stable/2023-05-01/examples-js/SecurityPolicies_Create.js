const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new security policy within the specified profile.
 *
 * @summary Creates a new security policy within the specified profile.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/SecurityPolicies_Create.json
 */
async function securityPoliciesCreate() {
  const subscriptionId = process.env["CDN_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["CDN_RESOURCE_GROUP"] || "RG";
  const profileName = "profile1";
  const securityPolicyName = "securityPolicy1";
  const securityPolicy = {
    parameters: {
      type: "WebApplicationFirewall",
      associations: [
        {
          domains: [
            {
              id: "/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/customdomains/testdomain1",
            },
            {
              id: "/subscriptions/subid/resourcegroups/RG/providers/Microsoft.Cdn/profiles/profile1/customdomains/testdomain2",
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
  const result = await client.securityPolicies.beginCreateAndWait(
    resourceGroupName,
    profileName,
    securityPolicyName,
    securityPolicy
  );
  console.log(result);
}
