/** Samples for L3Networks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/L3Networks_Delete.json
     */
    /**
     * Sample code: Delete L3 network.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deleteL3Network(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.l3Networks().delete("resourceGroupName", "l3NetworkName", com.azure.core.util.Context.NONE);
    }
}
