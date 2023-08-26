/** Samples for L3Networks ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2023-07-01/examples/L3Networks_ListByResourceGroup.json
     */
    /**
     * Sample code: List L3 networks for resource group.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void listL3NetworksForResourceGroup(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.l3Networks().listByResourceGroup("resourceGroupName", com.azure.core.util.Context.NONE);
    }
}
