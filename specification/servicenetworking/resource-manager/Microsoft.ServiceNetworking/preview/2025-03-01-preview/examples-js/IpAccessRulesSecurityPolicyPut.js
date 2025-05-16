const { ServiceNetworkingManagementClient } = require("@azure/arm-servicenetworking");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a SecurityPolicy
 *
 * @summary create a SecurityPolicy
 * x-ms-original-file: 2025-03-01-preview/IpAccessRulesSecurityPolicyPut.json
 */
async function putIpAccessRulesSecurityPolicy() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subid";
  const client = new ServiceNetworkingManagementClient(credential, subscriptionId);
  const result = await client.securityPoliciesInterface.createOrUpdate("rg1", "tc1", "sp1", {
    location: "NorthCentralUS",
    properties: { ipAccessRulesPolicy: { rules: [] } },
  });
  console.log(result);
}
