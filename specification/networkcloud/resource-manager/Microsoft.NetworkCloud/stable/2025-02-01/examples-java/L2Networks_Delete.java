
/**
 * Samples for L2Networks Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/L2Networks_Delete.
     * json
     */
    /**
     * Sample code: Delete L2 network.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deleteL2Network(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.l2Networks().delete("resourceGroupName", "l2NetworkName", null, null, com.azure.core.util.Context.NONE);
    }
}
