const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to delete a gateway.
 *
 * @summary The operation to delete a gateway.
 * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-07-31-preview/examples/gateway/Gateway_Delete.json
 */
async function deleteAGateway() {
  const subscriptionId =
    process.env["HYBRIDCOMPUTE_SUBSCRIPTION_ID"] || "ffd506c8-3415-42d3-9612-fdb423fb17df";
  const resourceGroupName = process.env["HYBRIDCOMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const gatewayName = "{gatewayName}";
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const result = await client.gateways.beginDeleteAndWait(resourceGroupName, gatewayName);
  console.log(result);
}
