
/**
 * Samples for StorageDiscoveryWorkspaces ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/StorageDiscoveryWorkspaces_ListByResourceGroup.json
     */
    /**
     * Sample code: List StorageDiscoveryWorkspaces by Resource Group.
     * 
     * @param manager Entry point to StorageDiscoveryManager.
     */
    public static void listStorageDiscoveryWorkspacesByResourceGroup(
        com.azure.resourcemanager.storagediscovery.StorageDiscoveryManager manager) {
        manager.storageDiscoveryWorkspaces().listByResourceGroup("sample-rg", com.azure.core.util.Context.NONE);
    }
}
