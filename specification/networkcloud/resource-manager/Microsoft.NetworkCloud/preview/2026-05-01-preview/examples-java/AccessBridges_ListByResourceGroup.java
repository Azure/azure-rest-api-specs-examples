
/**
 * Samples for AccessBridges ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/AccessBridges_ListByResourceGroup.json
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
