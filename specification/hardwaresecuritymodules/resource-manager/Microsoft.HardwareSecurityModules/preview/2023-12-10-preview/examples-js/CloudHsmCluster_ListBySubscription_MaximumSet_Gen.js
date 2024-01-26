const { AzureHSMResourceProvider } = require("@azure/arm-hardwaresecuritymodules");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The List operation gets information about the Cloud HSM Clusters associated with the subscription.
 *
 * @summary The List operation gets information about the Cloud HSM Clusters associated with the subscription.
 * x-ms-original-file: specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2023-12-10-preview/examples/CloudHsmCluster_ListBySubscription_MaximumSet_Gen.json
 */
async function cloudHsmClusterListBySubscriptionMaximumSetGen() {
  const subscriptionId =
    process.env["HARDWARESECURITYMODULES_SUBSCRIPTION_ID"] ||
    "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new AzureHSMResourceProvider(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.cloudHsmClusters.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
