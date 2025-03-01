const { ImpactClient } = require("@azure/arm-impactreporting");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a WorkloadImpact
 *
 * @summary create a WorkloadImpact
 * x-ms-original-file: 2024-05-01-preview/WorkloadPerformance_Create.json
 */
async function reportingPerformanceRelatedImpact() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ImpactClient(credential, subscriptionId);
  const result = await client.workloadImpacts.create("impact-002", {
    properties: {
      impactedResourceId:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-rg/providers/Microsoft.Sql/sqlserver/dbservercontext",
      startDateTime: new Date("2022-06-15T05:59:46.6517821Z"),
      impactDescription: "high cpu utilization",
      impactCategory: "Resource.Performance",
      workload: { context: "webapp/scenario1", toolset: "Other" },
      performance: [{ metricName: "CPU", actual: 90, expected: 60 }],
      clientIncidentDetails: {
        clientIncidentId: "AA123",
        clientIncidentSource: "Jira",
      },
    },
  });
  console.log(result);
}
