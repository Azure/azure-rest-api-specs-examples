
/**
 * Samples for StorageContainers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/StorageContainers_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: StorageContainers_Get_MaximumSet.
     * 
     * @param manager Entry point to DiscoveryManager.
     */
    public static void storageContainersGetMaximumSet(com.azure.resourcemanager.discovery.DiscoveryManager manager) {
        manager.storageContainers().getByResourceGroupWithResponse("rgdiscovery", "60fa9761e5831e6b1e",
            com.azure.core.util.Context.NONE);
    }
}
