
/**
 * Samples for ArcSettings Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * GetArcSetting.json
     */
    /**
     * Sample code: Get ArcSetting.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void getArcSetting(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.arcSettings().getWithResponse("test-rg", "myCluster", "default", com.azure.core.util.Context.NONE);
    }
}
