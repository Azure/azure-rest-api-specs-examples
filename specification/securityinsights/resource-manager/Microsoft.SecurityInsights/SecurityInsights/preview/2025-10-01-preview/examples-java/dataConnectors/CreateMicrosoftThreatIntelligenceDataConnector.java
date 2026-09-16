
import com.azure.resourcemanager.securityinsights.models.DataTypeState;
import com.azure.resourcemanager.securityinsights.models.MSTIDataConnector;
import com.azure.resourcemanager.securityinsights.models.MSTIDataConnectorDataTypes;
import com.azure.resourcemanager.securityinsights.models.MstiDataConnectorDataTypesMicrosoftEmergingThreatFeed;
import java.time.OffsetDateTime;

/**
 * Samples for DataConnectors CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/dataConnectors/CreateMicrosoftThreatIntelligenceDataConnector.json
     */
    /**
     * Sample code: Creates or updates a Microsoft Threat Intelligence data connector.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createsOrUpdatesAMicrosoftThreatIntelligenceDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectors().createOrUpdateWithResponse("myRg", "myWorkspace",
            "c345bf40-8509-4ed2-b947-50cb773aaf04",
            new MSTIDataConnector()
                .withDataTypes(new MSTIDataConnectorDataTypes().withMicrosoftEmergingThreatFeed(
                    new MstiDataConnectorDataTypesMicrosoftEmergingThreatFeed().withState(DataTypeState.ENABLED)
                        .withLookbackPeriod(OffsetDateTime.parse("1970-01-01T00:00:00.000Z"))))
                .withTenantId("06b3ccb8-1384-4bcc-aec7-852f6d57161b"),
            com.azure.core.util.Context.NONE);
    }
}
