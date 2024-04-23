const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update an existing CdnWebApplicationFirewallPolicy with the specified policy name under the specified subscription and resource group
 *
 * @summary Update an existing CdnWebApplicationFirewallPolicy with the specified policy name under the specified subscription and resource group
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/WafPatchPolicy.json
 */
async function createsSpecificPolicy() {
  const subscriptionId = process.env["CDN_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["CDN_RESOURCE_GROUP"] || "rg1";
  const policyName = "MicrosoftCdnWafPolicy";
  const cdnWebApplicationFirewallPolicyPatchParameters = { tags: { foo: "bar" } };
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.policies.beginUpdateAndWait(
    resourceGroupName,
    policyName,
    cdnWebApplicationFirewallPolicyPatchParameters,
  );
  console.log(result);
}
