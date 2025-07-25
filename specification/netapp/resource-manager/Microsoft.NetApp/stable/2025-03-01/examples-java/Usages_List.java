
/**
 * Samples for NetAppResourceUsages List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-03-01/examples/Usages_List.json
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
