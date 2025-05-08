
/**
 * Samples for Pools Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-01-01-preview/examples/
     * Pools_Get_CustomThroughput.json
     */
    /**
     * Sample code: Pools_Get_CustomThroughput.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void poolsGetCustomThroughput(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.pools().getWithResponse("myRG", "account1", "customPool1", com.azure.core.util.Context.NONE);
    }
}
