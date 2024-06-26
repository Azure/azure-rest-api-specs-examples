const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a Azure Arc PrivateLinkScope's validation details for a given machine.
 *
 * @summary Returns a Azure Arc PrivateLinkScope's validation details for a given machine.
 * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2021-06-10-preview/examples/PrivateLinkScopesGetValidationForMachine.json
 */
async function privateLinkScopeGet() {
  const subscriptionId = "86dc51d3-92ed-4d7e-947a-775ea79b4919";
  const resourceGroupName = "my-resource-group";
  const machineName = "machineName";
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const result = await client.privateLinkScopes.getValidationDetailsForMachine(
    resourceGroupName,
    machineName
  );
  console.log(result);
}

privateLinkScopeGet().catch(console.error);
