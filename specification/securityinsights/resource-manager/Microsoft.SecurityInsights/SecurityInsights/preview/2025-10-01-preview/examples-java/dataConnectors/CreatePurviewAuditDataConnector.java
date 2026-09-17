
import com.azure.resourcemanager.securityinsights.models.DCRConfiguration;
import com.azure.resourcemanager.securityinsights.models.DataTypeState;
import com.azure.resourcemanager.securityinsights.models.PurviewAuditConnectorDataTypes;
import com.azure.resourcemanager.securityinsights.models.PurviewAuditConnectorDataTypesLogs;
import com.azure.resourcemanager.securityinsights.models.PurviewAuditDataConnector;

/**
 * Samples for DataConnectors CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/dataConnectors/CreatePurviewAuditDataConnector.json
     */
    /**
     * Sample code: Creates or updates a PurviewAudit data connector.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createsOrUpdatesAPurviewAuditDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectors()
            .createOrUpdateWithResponse("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
                new PurviewAuditDataConnector().withEtag("\"0300bf09-0000-0000-0000-5c37296e0000\"")
                    .withConnectorDefinitionName("PowerAutomate").withSourceType("MicrosoftFlow")
                    .withDcrConfig(new DCRConfiguration().withDataCollectionEndpoint(
                        "https://microsoft-sentinel-datacollectionendpoint-123m.westeurope-1.ingest.monitor.azure.com")
                        .withDataCollectionRuleImmutableId("dcr-de21b053bd5a44beb99a256c9db85023")
                        .withStreamName("OFFICEPOWERAUTOMATE_RESTAPI"))
                    .withDataTypes(new PurviewAuditConnectorDataTypes()
                        .withLogs(new PurviewAuditConnectorDataTypesLogs().withState(DataTypeState.ENABLED)))
                    .withTenantId("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
                com.azure.core.util.Context.NONE);
    }
}
