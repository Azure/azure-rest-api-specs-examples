const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Policies for protecting resources using Just-in-Time access control for the subscription, location
 *
 * @summary Policies for protecting resources using Just-in-Time access control for the subscription, location
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/JitNetworkAccessPolicies/GetJitNetworkAccessPoliciesSubscriptionLocation_example.json
 */
async function getJitNetworkAccessPoliciesOnASubscriptionFromASecurityDataLocation() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const ascLocation = "westeurope";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.jitNetworkAccessPolicies.listByRegion(ascLocation)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getJitNetworkAccessPoliciesOnASubscriptionFromASecurityDataLocation().catch(console.error);
