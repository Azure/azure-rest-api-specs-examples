
/**
 * Samples for NetAppResourceUsages Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-06-01/examples/Usages_Get.json
     */
    /**
     * Sample code: Usages_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void usagesGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.netAppResourceUsages().getWithResponse("eastus", "totalTibsPerSubscription",
            com.azure.core.util.Context.NONE);
    }
}
