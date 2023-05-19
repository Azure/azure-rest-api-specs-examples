/** Samples for HybridAksClusters Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/HybridAksClusters_Delete.json
     */
    /**
     * Sample code: Delete Hybrid AKS provisioned cluster data.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deleteHybridAKSProvisionedClusterData(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager
            .hybridAksClusters()
            .delete("resourceGroupName", "hybridAksClusterName", com.azure.core.util.Context.NONE);
    }
}
