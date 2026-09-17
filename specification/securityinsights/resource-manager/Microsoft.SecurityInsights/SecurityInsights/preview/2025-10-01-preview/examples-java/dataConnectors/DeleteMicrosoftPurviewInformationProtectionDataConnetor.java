
/**
 * Samples for DataConnectors Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-10-01-preview/dataConnectors/DeleteMicrosoftPurviewInformationProtectionDataConnetor.json
     */
    /**
     * Sample code: Delete an MicrosoftPurviewInformationProtection data connector.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void deleteAnMicrosoftPurviewInformationProtectionDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectors().deleteWithResponse("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
            com.azure.core.util.Context.NONE);
    }
}
