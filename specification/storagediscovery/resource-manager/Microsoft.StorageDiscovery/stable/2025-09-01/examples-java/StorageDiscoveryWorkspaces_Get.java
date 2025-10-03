
/**
 * Samples for StorageDiscoveryWorkspaces GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/StorageDiscoveryWorkspaces_Get.json
     */
    /**
     * Sample code: Get a StorageDiscoveryWorkspace.
     * 
     * @param manager Entry point to StorageDiscoveryManager.
     */
    public static void
        getAStorageDiscoveryWorkspace(com.azure.resourcemanager.storagediscovery.StorageDiscoveryManager manager) {
        manager.storageDiscoveryWorkspaces().getByResourceGroupWithResponse("sample-rg", "Sample-Storage-Workspace",
            com.azure.core.util.Context.NONE);
    }
}
