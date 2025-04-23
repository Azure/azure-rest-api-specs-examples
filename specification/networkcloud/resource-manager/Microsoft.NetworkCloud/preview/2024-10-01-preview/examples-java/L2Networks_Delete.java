
/**
 * Samples for L2Networks Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2024-10-01-preview/examples/
     * L2Networks_Delete.json
     */
    /**
     * Sample code: Delete L2 network.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deleteL2Network(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.l2Networks().delete("resourceGroupName", "l2NetworkName", com.azure.core.util.Context.NONE);
    }
}
