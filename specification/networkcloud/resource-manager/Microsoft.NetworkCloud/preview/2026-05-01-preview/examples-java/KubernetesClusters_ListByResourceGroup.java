
/**
 * Samples for KubernetesClusters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/KubernetesClusters_ListByResourceGroup.json
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
