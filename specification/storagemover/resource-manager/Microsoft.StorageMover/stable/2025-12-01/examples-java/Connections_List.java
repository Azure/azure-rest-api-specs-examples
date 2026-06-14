
/**
 * Samples for Connections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/Connections_List.json
     */
    /**
     * Sample code: Connections_List.
     * 
     * @param manager Entry point to StorageMoverManager.
     */
    public static void connectionsList(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager.connections().list("examples-rg", "examples-storageMoverName", com.azure.core.util.Context.NONE);
    }
}
