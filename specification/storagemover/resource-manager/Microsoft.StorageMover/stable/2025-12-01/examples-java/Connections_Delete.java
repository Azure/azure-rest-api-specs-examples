
/**
 * Samples for Connections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/Connections_Delete.json
     */
    /**
     * Sample code: Connections_Delete.
     * 
     * @param manager Entry point to StorageMoverManager.
     */
    public static void connectionsDelete(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager.connections().delete("examples-rg", "examples-storageMoverName", "example-connection",
            com.azure.core.util.Context.NONE);
    }
}
