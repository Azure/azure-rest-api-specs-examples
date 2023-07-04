/** Samples for BmcKeySets Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/BmcKeySets_Delete.json
     */
    /**
     * Sample code: Delete baseboard management controller key set of cluster.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deleteBaseboardManagementControllerKeySetOfCluster(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager
            .bmcKeySets()
            .delete("resourceGroupName", "clusterName", "bmcKeySetName", com.azure.core.util.Context.NONE);
    }
}
