const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the current status of IDPS signatures for the relevant policy
 *
 * @summary Retrieves the current status of IDPS signatures for the relevant policy
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/FirewallPolicyQuerySignatureOverrides.json
 */
async function querySignatureOverrides() {
  const subscriptionId =
    process.env["NETWORK_SUBSCRIPTION_ID"] || "e747cc13-97d4-4a79-b463-42d7f4e558f2";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const firewallPolicyName = "firewallPolicy";
  const parameters = {
    filters: [{ field: "Mode", values: ["Deny"] }],
    orderBy: { field: "severity", order: "Ascending" },
    resultsPerPage: 20,
    search: "",
    skip: 0,
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.firewallPolicyIdpsSignatures.list(
    resourceGroupName,
    firewallPolicyName,
    parameters
  );
  console.log(result);
}
