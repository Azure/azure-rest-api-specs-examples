
import com.azure.resourcemanager.securityinsights.models.DataTypeState;
import com.azure.resourcemanager.securityinsights.models.PremiumMdtiDataConnectorDataTypes;
import com.azure.resourcemanager.securityinsights.models.PremiumMdtiDataConnectorDataTypesConnector;
import com.azure.resourcemanager.securityinsights.models.PremiumMicrosoftDefenderForThreatIntelligence;
import java.time.OffsetDateTime;

/**
 * Samples for DataConnectors CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-10-01-preview/dataConnectors/CreatePremiumMicrosoftDefenderForThreatIntelligenceDataConnector.json
     */
    /**
     * Sample code: Creates or updates a PremiumMicrosoftDefenderForThreatIntelligence data connector.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createsOrUpdatesAPremiumMicrosoftDefenderForThreatIntelligenceDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectors().createOrUpdateWithResponse("myRg", "myWorkspace",
            "8c569548-a86c-4fb4-8ae4-d1e35a6146f8",
            new PremiumMicrosoftDefenderForThreatIntelligence()
                .withLookbackPeriod(OffsetDateTime.parse("1970-01-01T00:00:00.000Z"))
                .withDataTypes(new PremiumMdtiDataConnectorDataTypes()
                    .withConnector(new PremiumMdtiDataConnectorDataTypesConnector().withState(DataTypeState.ENABLED)))
                .withTenantId("e4afb3c4-813b-4e68-b6de-e5360866e798"),
            com.azure.core.util.Context.NONE);
    }
}
