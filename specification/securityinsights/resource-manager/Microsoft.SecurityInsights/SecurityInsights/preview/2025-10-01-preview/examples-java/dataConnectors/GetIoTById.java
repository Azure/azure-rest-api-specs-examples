
/**
 * Samples for DataConnectors Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/dataConnectors/GetIoTById.json
     */
    /**
     * Sample code: Get a IoT data connector.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getAIoTDataConnector(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.dataConnectors().getWithResponse("myRg", "myWorkspace", "d2e5dc7a-f3a2-429d-954b-939fa8c2932e",
            com.azure.core.util.Context.NONE);
    }
}
