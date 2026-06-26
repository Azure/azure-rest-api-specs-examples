
/**
 * Samples for KubernetesVersions ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/KubernetesVersions_ListByResourceGroup.json
     */
    /**
     * Sample code: List Kubernetes versions for resource group.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listKubernetesVersionsForResourceGroup(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.kubernetesVersions().listByResourceGroup("resourceGroupName", null, null,
            com.azure.core.util.Context.NONE);
    }
}
