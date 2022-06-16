const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Policies for protecting resources using Just-in-Time access control for the subscription, location
 *
 * @summary Policies for protecting resources using Just-in-Time access control for the subscription, location
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/GetJitNetworkAccessPolicy_example.json
 */
async function getJitNetworkAccessPolicy() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const resourceGroupName = "myRg1";
  const ascLocation = "westeurope";
  const jitNetworkAccessPolicyName = "default";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.jitNetworkAccessPolicies.get(
    resourceGroupName,
    ascLocation,
    jitNetworkAccessPolicyName
  );
  console.log(result);
}

getJitNetworkAccessPolicy().catch(console.error);
