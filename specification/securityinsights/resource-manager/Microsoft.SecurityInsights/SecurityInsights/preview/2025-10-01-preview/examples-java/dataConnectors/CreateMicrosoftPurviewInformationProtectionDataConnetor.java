
import com.azure.resourcemanager.securityinsights.models.DataTypeState;
import com.azure.resourcemanager.securityinsights.models.MicrosoftPurviewInformationProtectionConnectorDataTypes;
import com.azure.resourcemanager.securityinsights.models.MicrosoftPurviewInformationProtectionConnectorDataTypesLogs;
import com.azure.resourcemanager.securityinsights.models.MicrosoftPurviewInformationProtectionDataConnector;

/**
 * Samples for DataConnectors CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-10-01-preview/dataConnectors/CreateMicrosoftPurviewInformationProtectionDataConnetor.json
     */
    /**
     * Sample code: Creates or updates an MicrosoftPurviewInformationProtection data connector.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createsOrUpdatesAnMicrosoftPurviewInformationProtectionDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectors().createOrUpdateWithResponse("myRg", "myWorkspace",
            "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
            new MicrosoftPurviewInformationProtectionDataConnector()
                .withEtag("\"0300bf09-0000-0000-0000-5c37296e0000\"")
                .withDataTypes(new MicrosoftPurviewInformationProtectionConnectorDataTypes().withLogs(
                    new MicrosoftPurviewInformationProtectionConnectorDataTypesLogs().withState(DataTypeState.ENABLED)))
                .withTenantId("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
            com.azure.core.util.Context.NONE);
    }
}
