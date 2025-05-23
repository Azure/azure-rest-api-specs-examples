const { AzureHSMResourceProvider } = require("@azure/arm-hardwaresecuritymodules");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a Cloud HSM Cluster in the specified subscription.
 *
 * @summary Update a Cloud HSM Cluster in the specified subscription.
 * x-ms-original-file: specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2024-06-30-preview/examples/CloudHsmCluster_Update_MaximumSet_Gen.json
 */
async function cloudHsmClusterUpdateMaximumSetGen() {
  const subscriptionId =
    process.env["HARDWARESECURITYMODULES_SUBSCRIPTION_ID"] ||
    "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["HARDWARESECURITYMODULES_RESOURCE_GROUP"] || "rgcloudhsm";
  const cloudHsmClusterName = "chsm1";
  const tags = { dept: "hsm", environment: "dogfood", slice: "A" };
  const identity = {
    type: "UserAssigned",
    userAssignedIdentities: {
      "/subscriptions/00000000000000000000000000000000/resourceGroups/contosoResources/providers/MicrosoftManagedIdentity/userAssignedIdentities/identity1":
        {},
    },
  };
  const options = { tags, identity };
  const credential = new DefaultAzureCredential();
  const client = new AzureHSMResourceProvider(credential, subscriptionId);
  const result = await client.cloudHsmClusters.beginUpdateAndWait(
    resourceGroupName,
    cloudHsmClusterName,
    options,
  );
  console.log(result);
}
