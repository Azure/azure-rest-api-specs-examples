const { AzureHSMResourceProvider } = require("@azure/arm-hardwaresecuritymodules");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the private endpoint connection for the Cloud Hsm Cluster.
 *
 * @summary Gets the private endpoint connection for the Cloud Hsm Cluster.
 * x-ms-original-file: specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2023-12-10-preview/examples/CloudHsmClusterPrivateEndpointConnection_Get_MaximumSet_Gen.json
 */
async function cloudHsmClusterPrivateEndpointConnectionGetMaximumSetGen() {
  const subscriptionId =
    process.env["HARDWARESECURITYMODULES_SUBSCRIPTION_ID"] ||
    "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["HARDWARESECURITYMODULES_RESOURCE_GROUP"] || "rgcloudhsm";
  const cloudHsmClusterName = "chsm1";
  const peConnectionName = "sample-pec";
  const credential = new DefaultAzureCredential();
  const client = new AzureHSMResourceProvider(credential, subscriptionId);
  const result = await client.cloudHsmClusterPrivateEndpointConnections.get(
    resourceGroupName,
    cloudHsmClusterName,
    peConnectionName,
  );
  console.log(result);
}
