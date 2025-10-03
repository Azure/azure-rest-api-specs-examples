
/**
 * Samples for StorageDiscoveryWorkspaces Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/StorageDiscoveryWorkspaces_Delete.json
     */
    /**
     * Sample code: Delete a StorageDiscoveryWorkspace.
     * 
     * @param manager Entry point to StorageDiscoveryManager.
     */
    public static void
        deleteAStorageDiscoveryWorkspace(com.azure.resourcemanager.storagediscovery.StorageDiscoveryManager manager) {
        manager.storageDiscoveryWorkspaces().deleteByResourceGroupWithResponse("sample-rg", "sampleworkspace",
            com.azure.core.util.Context.NONE);
    }
}
