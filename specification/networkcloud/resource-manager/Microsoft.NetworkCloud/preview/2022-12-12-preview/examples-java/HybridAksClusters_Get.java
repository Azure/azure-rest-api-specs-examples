/** Samples for HybridAksClusters GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/HybridAksClusters_Get.json
     */
    /**
     * Sample code: Get Hybrid AKS provisioned cluster data.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void getHybridAKSProvisionedClusterData(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager
            .hybridAksClusters()
            .getByResourceGroupWithResponse(
                "resourceGroupName", "hybridAksClusterName", com.azure.core.util.Context.NONE);
    }
}
