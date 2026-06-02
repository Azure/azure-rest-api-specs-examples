const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a compute associated with the Cognitive Services account.
 *
 * @summary creates or updates a compute associated with the Cognitive Services account.
 * x-ms-original-file: 2026-03-15-preview/PutCompute.json
 */
async function putCompute() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.computes.createOrUpdate(
    "rgcognitiveservices",
    "myAccount",
    "myCompute",
    {
      properties: {
        computeType: "Cluster",
        pools: [
          { name: "default", vmPriority: "Regular", instanceType: "Standard_DS3_v2", nodeCount: 2 },
        ],
        subnetArmId:
          "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rgcognitiveservices/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/default",
      },
      location: "eastus",
      identity: { type: "None" },
    },
  );
  console.log(result);
}
