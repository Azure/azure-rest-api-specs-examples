const { ImpactClient } = require("@azure/arm-impactreporting");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a WorkloadImpact
 *
 * @summary create a WorkloadImpact
 * x-ms-original-file: 2024-05-01-preview/WorkloadConnectivityImpact_Create.json
 */
async function reportingAConnectivityImpact() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ImpactClient(credential, subscriptionId);
  const result = await client.workloadImpacts.create("impact-001", {
    properties: {
      impactedResourceId:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-rg/providers/Microsoft.Sql/sqlserver/dbservercontext",
      startDateTime: new Date("2022-06-15T05:59:46.6517821Z"),
      impactDescription: "conection failure",
      impactCategory: "Resource.Connectivity",
      connectivity: {
        protocol: "TCP",
        port: 1443,
        source: {
          azureResourceId:
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceSub/providers/Microsoft.compute/virtualmachines/vm1",
        },
        target: {
          azureResourceId:
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceSub/providers/Microsoft.compute/virtualmachines/vm2",
        },
      },
      workload: { context: "webapp/scenario1", toolset: "Other" },
      clientIncidentDetails: {
        clientIncidentId: "AA123",
        clientIncidentSource: "Jira",
      },
    },
  });
  console.log(result);
}
