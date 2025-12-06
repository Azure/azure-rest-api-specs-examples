
/**
 * Samples for Pools List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Pools_List.json
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
