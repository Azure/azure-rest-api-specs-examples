const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Will override/create a new signature overrides for the policy's IDPS
 *
 * @summary Will override/create a new signature overrides for the policy's IDPS
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/FirewallPolicySignatureOverridesPut.json
 */
async function putSignatureOverrides() {
  const subscriptionId = "e747cc13-97d4-4a79-b463-42d7f4e558f2";
  const resourceGroupName = "rg1";
  const firewallPolicyName = "firewallPolicy";
  const parameters = {
    name: "default",
    type: "Microsoft.Network/firewallPolicies/signatureOverrides",
    id: "/subscriptions/e747cc13-97d4-4a79-b463-42d7f4e558f2/resourceGroups/rg1/providers/Microsoft.Network/firewallPolicies/firewallPolicy/signatureOverrides/default",
    properties: { signatures: { "2000105": "Off", "2000106": "Deny" } },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.firewallPolicyIdpsSignaturesOverrides.put(
    resourceGroupName,
    firewallPolicyName,
    parameters
  );
  console.log(result);
}

putSignatureOverrides().catch(console.error);
