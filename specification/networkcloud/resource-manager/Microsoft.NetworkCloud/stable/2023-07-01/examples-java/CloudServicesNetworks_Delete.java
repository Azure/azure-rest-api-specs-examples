/** Samples for CloudServicesNetworks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2023-07-01/examples/CloudServicesNetworks_Delete.json
     */
    /**
     * Sample code: Delete cloud services network.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deleteCloudServicesNetwork(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager
            .cloudServicesNetworks()
            .delete("resourceGroupName", "cloudServicesNetworkName", com.azure.core.util.Context.NONE);
    }
}
