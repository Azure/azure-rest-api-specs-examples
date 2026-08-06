
/**
 * Samples for StorageContainers ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/StorageContainers_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: StorageContainers_ListByResourceGroup_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void
        storageContainersListByResourceGroupMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.storageContainers().listByResourceGroup("rgdiscovery", com.azure.core.util.Context.NONE);
    }
}
