/** Samples for CloudServicesNetworks GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/CloudServicesNetworks_Get.json
     */
    /**
     * Sample code: Get cloud services network.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void getCloudServicesNetwork(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager
            .cloudServicesNetworks()
            .getByResourceGroupWithResponse(
                "resourceGroupName", "cloudServicesNetworkName", com.azure.core.util.Context.NONE);
    }
}
