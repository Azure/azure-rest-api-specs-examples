
/**
 * Samples for Pools List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-01-01-preview/examples/Pools_List.json
     */
    /**
     * Sample code: Pools_List.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void poolsList(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.pools().list("myRG", "account1", com.azure.core.util.Context.NONE);
    }
}
