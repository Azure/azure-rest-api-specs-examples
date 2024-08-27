
/**
 * Samples for Extensions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * GetExtension.json
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
