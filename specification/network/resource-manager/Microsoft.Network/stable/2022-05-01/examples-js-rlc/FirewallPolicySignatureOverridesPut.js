const createNetworkManagementClient = require("@azure-rest/arm-network").default;
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Will override/create a new signature overrides for the policy's IDPS
 *
 * @summary Will override/create a new signature overrides for the policy's IDPS
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/FirewallPolicySignatureOverridesPut.json
 */
async function putSignatureOverrides() {
  const credential = new DefaultAzureCredential();
  const client = createNetworkManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "rg1";
  const firewallPolicyName = "firewallPolicy";
  const options = {
    body: {
      name: "default",
      type: "Microsoft.Network/firewallPolicies/signatureOverrides",
      id: "/subscriptions/e747cc13-97d4-4a79-b463-42d7f4e558f2/resourceGroups/rg1/providers/Microsoft.Network/firewallPolicies/firewallPolicy/signatureOverrides/default",
      properties: { signatures: { 2000105: "Off", 2000106: "Deny" } },
    },
    queryParameters: { "api-version": "2022-05-01" },
  };
  const result = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/signatureOverrides/default",
      subscriptionId,
      resourceGroupName,
      firewallPolicyName,
    )
    .put(options);
  console.log(result);
}

putSignatureOverrides().catch(console.error);
