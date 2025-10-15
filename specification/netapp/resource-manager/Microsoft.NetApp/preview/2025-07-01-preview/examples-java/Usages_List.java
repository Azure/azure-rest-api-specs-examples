
/**
 * Samples for NetAppResourceUsages List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-07-01-preview/examples/Usages_List.json
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
