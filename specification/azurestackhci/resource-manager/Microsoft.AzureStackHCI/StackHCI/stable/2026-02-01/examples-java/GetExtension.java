
/**
 * Samples for Extensions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/GetExtension.json
     */
    /**
     * Sample code: Get ArcSettings Extension.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void getArcSettingsExtension(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.extensions().getWithResponse("test-rg", "myCluster", "default", "MicrosoftMonitoringAgent",
            com.azure.core.util.Context.NONE);
    }
}
