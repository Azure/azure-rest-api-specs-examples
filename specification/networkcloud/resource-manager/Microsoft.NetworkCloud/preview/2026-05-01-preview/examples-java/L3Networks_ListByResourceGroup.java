
/**
 * Samples for L3Networks ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/L3Networks_ListByResourceGroup.json
     */
    /**
     * Sample code: List L3 networks for resource group.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listL3NetworksForResourceGroup(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.l3Networks().listByResourceGroup("resourceGroupName", null, null, com.azure.core.util.Context.NONE);
    }
}
