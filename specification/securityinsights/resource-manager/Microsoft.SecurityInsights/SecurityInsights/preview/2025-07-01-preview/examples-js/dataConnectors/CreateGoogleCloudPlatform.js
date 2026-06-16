const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates the data connector.
 *
 * @summary creates or updates the data connector.
 * x-ms-original-file: 2025-07-01-preview/dataConnectors/CreateGoogleCloudPlatform.json
 */
async function createsOrUpdatesAGCPDataConnector() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.dataConnectors.createOrUpdate(
    "myRg",
    "myWorkspace",
    "GCP_fce27b90-d6f5-4d30-991a-af509a2b50a1",
    {
      kind: "GCP",
      auth: {
        projectNumber: "123456789012",
        serviceAccountEmail: "sentinel-service-account@project-id.iam.gserviceaccount.com",
        workloadIdentityProviderId: "sentinel-identity-provider",
      },
      connectorDefinitionName: "GcpConnector",
      dcrConfig: {
        dataCollectionEndpoint:
          "https://microsoft-sentinel-datacollectionendpoint-123m.westeurope-1.ingest.monitor.azure.com",
        dataCollectionRuleImmutableId: "dcr-de21b053bd5a44beb99a256c9db85023",
        streamName: "SENTINEL_GCP_AUDIT_LOGS",
      },
      request: { projectId: "project-id", subscriptionNames: ["sentinel-subscription"] },
    },
  );
  console.log(result);
}
