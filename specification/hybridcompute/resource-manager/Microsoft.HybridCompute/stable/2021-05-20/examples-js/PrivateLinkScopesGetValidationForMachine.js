const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

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
