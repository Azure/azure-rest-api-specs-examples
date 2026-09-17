
import com.azure.resourcemanager.securityinsights.models.DCRConfiguration;
import com.azure.resourcemanager.securityinsights.models.GCPAuthProperties;
import com.azure.resourcemanager.securityinsights.models.GCPDataConnector;
import com.azure.resourcemanager.securityinsights.models.GCPRequestProperties;
import java.util.Arrays;

/**
 * Samples for DataConnectors CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/dataConnectors/CreateGoogleCloudPlatform.json
     */
    /**
     * Sample code: Creates or updates a GCP data connector.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        createsOrUpdatesAGCPDataConnector(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectors().createOrUpdateWithResponse("myRg", "myWorkspace",
            "GCP_fce27b90-d6f5-4d30-991a-af509a2b50a1",
            new GCPDataConnector().withConnectorDefinitionName("GcpConnector")
                .withAuth(new GCPAuthProperties()
                    .withServiceAccountEmail("sentinel-service-account@project-id.iam.gserviceaccount.com")
                    .withProjectNumber("123456789012").withWorkloadIdentityProviderId("sentinel-identity-provider"))
                .withRequest(new GCPRequestProperties().withProjectId("project-id")
                    .withSubscriptionNames(Arrays.asList("sentinel-subscription")))
                .withDcrConfig(new DCRConfiguration()
                    .withDataCollectionEndpoint(
                        "https://microsoft-sentinel-datacollectionendpoint-123m.westeurope-1.ingest.monitor.azure.com")
                    .withDataCollectionRuleImmutableId("dcr-de21b053bd5a44beb99a256c9db85023")
                    .withStreamName("SENTINEL_GCP_AUDIT_LOGS")),
            com.azure.core.util.Context.NONE);
    }
}
