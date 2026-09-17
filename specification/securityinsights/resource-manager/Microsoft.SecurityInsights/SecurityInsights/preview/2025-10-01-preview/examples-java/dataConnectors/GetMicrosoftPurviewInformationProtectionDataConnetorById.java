
/**
 * Samples for DataConnectors Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-10-01-preview/dataConnectors/GetMicrosoftPurviewInformationProtectionDataConnetorById.json
     */
    /**
     * Sample code: Get a MicrosoftPurviewInformationProtection data connector.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAMicrosoftPurviewInformationProtectionDataConnector(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectors().getWithResponse("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
            com.azure.core.util.Context.NONE);
    }
}
