const { WorkloadManagerClient } = require("@azure/arm-computeworkloadmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or replaces a runtime link.
 *
 * @summary creates or replaces a runtime link.
 * x-ms-original-file: 2026-11-01-preview/RuntimeLinks_CreateOrUpdate_AksAci.json
 */
async function createARuntimeLinkBetweenKubernetesAndServerlessContainers() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new WorkloadManagerClient(credential, subscriptionId);
  const result = await client.runtimeLinks.createOrUpdate(
    "rg-workload",
    "managed-agents-prod",
    "default",
    {
      location: "eastus2",
      properties: {
        orchestratorBindingResourceId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-workload/providers/Microsoft.Compute/workloadSpaces/managed-agents-prod/runtimeBindings/kubernetes-default",
        executionBindingResourceId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-workload/providers/Microsoft.Compute/workloadSpaces/managed-agents-prod/runtimeBindings/serverless-default",
        capacityProfile: { minimumNodes: 1, maximumNodes: 100 },
      },
    },
  );
  console.log(result);
}
