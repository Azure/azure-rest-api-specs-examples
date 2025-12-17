
/**
 * Samples for KubernetesClusters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/
     * KubernetesClusters_ListByResourceGroup.json
     */
    /**
     * Sample code: List Kubernetes clusters for resource group.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listKubernetesClustersForResourceGroup(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.kubernetesClusters().listByResourceGroup("resourceGroupName", null, null,
            com.azure.core.util.Context.NONE);
    }
}
