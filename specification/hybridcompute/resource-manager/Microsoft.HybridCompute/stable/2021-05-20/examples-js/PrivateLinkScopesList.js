const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

async function privateLinkScopesListJson() {
  const subscriptionId = "86dc51d3-92ed-4d7e-947a-775ea79b4919";
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateLinkScopes.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

privateLinkScopesListJson().catch(console.error);
