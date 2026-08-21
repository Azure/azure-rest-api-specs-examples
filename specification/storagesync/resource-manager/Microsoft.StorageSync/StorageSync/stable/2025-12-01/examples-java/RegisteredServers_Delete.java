
/**
 * Samples for RegisteredServers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/RegisteredServers_Delete.json
     */
    /**
     * Sample code: RegisteredServers_Delete.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void registeredServersDelete(com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        manager.registeredServers().delete("SampleResourceGroup_1", "SampleStorageSyncService_1",
            "41166691-ab03-43e9-ab3e-0330eda162ac", com.azure.core.util.Context.NONE);
    }
}
