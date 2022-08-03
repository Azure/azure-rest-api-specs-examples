const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about the specified SIM policy.
 *
 * @summary Gets information about the specified SIM policy.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/SimPolicyGet.json
 */
async function getSimPolicy() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const mobileNetworkName = "testMobileNetwork";
  const simPolicyName = "testPolicy";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.simPolicies.get(resourceGroupName, mobileNetworkName, simPolicyName);
  console.log(result);
}

getSimPolicy().catch(console.error);
