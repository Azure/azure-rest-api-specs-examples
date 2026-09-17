
/**
 * Samples for AccessBridges ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/AccessBridges_ListByResourceGroup.json
     */
    /**
     * Sample code: List access bridges for resource group.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listAccessBridgesForResourceGroup(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.accessBridges().listByResourceGroup("resourceGroupName", null, null, com.azure.core.util.Context.NONE);
    }
}
