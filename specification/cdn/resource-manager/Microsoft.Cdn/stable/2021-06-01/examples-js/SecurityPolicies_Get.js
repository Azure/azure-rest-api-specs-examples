const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an existing security policy within a profile.
 *
 * @summary Gets an existing security policy within a profile.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/SecurityPolicies_Get.json
 */
async function securityPoliciesGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const securityPolicyName = "securityPolicy1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.securityPolicies.get(
    resourceGroupName,
    profileName,
    securityPolicyName
  );
  console.log(result);
}

securityPoliciesGet().catch(console.error);
