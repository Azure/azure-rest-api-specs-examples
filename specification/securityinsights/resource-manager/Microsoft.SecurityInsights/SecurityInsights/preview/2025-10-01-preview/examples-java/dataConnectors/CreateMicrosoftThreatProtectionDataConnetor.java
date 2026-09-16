
import com.azure.resourcemanager.securityinsights.models.DataTypeState;
import com.azure.resourcemanager.securityinsights.models.MTPDataConnector;
import com.azure.resourcemanager.securityinsights.models.MTPDataConnectorDataTypes;
import com.azure.resourcemanager.securityinsights.models.MTPDataConnectorDataTypesAlerts;
import com.azure.resourcemanager.securityinsights.models.MTPDataConnectorDataTypesIncidents;
import com.azure.resourcemanager.securityinsights.models.MtpFilteredProviders;
import com.azure.resourcemanager.securityinsights.models.MtpProvider;
import java.util.Arrays;

/**
 * Samples for DataConnectors CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/dataConnectors/CreateMicrosoftThreatProtectionDataConnetor.json
     */
    /**
     * Sample code: Creates or updates a MicrosoftThreatProtection data connector.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createsOrUpdatesAMicrosoftThreatProtectionDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectors().createOrUpdateWithResponse("myRg", "myWorkspace",
            "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
            new MTPDataConnector().withEtag("\"0300bf09-0000-0000-0000-5c37296e0000\"")
                .withDataTypes(new MTPDataConnectorDataTypes()
                    .withIncidents(new MTPDataConnectorDataTypesIncidents().withState(DataTypeState.DISABLED))
                    .withAlerts(new MTPDataConnectorDataTypesAlerts().withState(DataTypeState.ENABLED)))
                .withFilteredProviders(
                    new MtpFilteredProviders().withAlerts(Arrays.asList(MtpProvider.MICROSOFT_DEFENDER_FOR_CLOUD_APPS)))
                .withTenantId("178265c4-3136-4ff6-8ed1-b5b62b4cb5f5"),
            com.azure.core.util.Context.NONE);
    }
}
