const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an existing security policy within profile.
 *
 * @summary Deletes an existing security policy within profile.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/SecurityPolicies_Delete.json
 */
async function securityPoliciesDelete() {
  const subscriptionId = process.env["CDN_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["CDN_RESOURCE_GROUP"] || "RG";
  const profileName = "profile1";
  const securityPolicyName = "securityPolicy1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.securityPolicies.beginDeleteAndWait(
    resourceGroupName,
    profileName,
    securityPolicyName,
  );
  console.log(result);
}
