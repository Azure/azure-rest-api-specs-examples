const { PrometheusRuleGroupsManagementClient } = require("@azure/arm-prometheusrulegroups");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Prometheus rule group definition.
 *
 * @summary create or update a Prometheus rule group definition.
 * x-ms-original-file: 2023-03-01/createOrUpdateClusterCentricRuleGroup.json
 */
async function createOrUpdateAClusterCentricPrometheusRuleGroup() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new PrometheusRuleGroupsManagementClient(credential, subscriptionId);
  const result = await client.prometheusRuleGroups.createOrUpdate(
    "promResourceGroup",
    "myPrometheusRuleGroup",
    {
      location: "East US",
      description: "This is a rule group with culster centric configuration",
      clusterName: "myClusterName",
      interval: "PT10M",
      rules: [
        {
          actions: [],
          alert: "Billing_Processing_Very_Slow",
          annotations: { annotationName1: "annotationValue1" },
          enabled: true,
          expression: "job_type:billing_jobs_duration_seconds:99p5m > 30",
          for: "PT5M",
          labels: { team: "prod" },
          resolveConfiguration: { autoResolved: true, timeToResolve: "PT10M" },
          severity: 2,
        },
      ],
      scopes: [
        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/myResourceGroup/providers/microsoft.monitor/accounts/myAzureMonitorWorkspace",
        "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/myResourceGroup/providers/Microsoft.ContainerService/managedClusters/myClusterName",
      ],
    },
  );
  console.log(result);
}
