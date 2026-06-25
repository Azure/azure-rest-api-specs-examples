
/**
 * Samples for ClusterManagers ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/ClusterManagers_ListByResourceGroup.json
     */
    /**
     * Sample code: List cluster managers for resource group.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listClusterManagersForResourceGroup(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.clusterManagers().listByResourceGroup("resourceGroupName", null, null,
            com.azure.core.util.Context.NONE);
    }
}
