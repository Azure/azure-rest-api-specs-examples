
/**
 * Samples for L2Networks ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2025-07-01-preview/examples/
     * L2Networks_ListByResourceGroup.json
     */
    /**
     * Sample code: List L2 networks for resource group.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listL2NetworksForResourceGroup(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.l2Networks().listByResourceGroup("resourceGroupName", null, null, com.azure.core.util.Context.NONE);
    }
}
