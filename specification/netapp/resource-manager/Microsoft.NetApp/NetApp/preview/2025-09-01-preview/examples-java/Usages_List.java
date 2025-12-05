
/**
 * Samples for NetAppResourceUsages List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Usages_List.json
     */
    /**
     * Sample code: Usages_List.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void usagesList(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.netAppResourceUsages().list("eastus", com.azure.core.util.Context.NONE);
    }
}
