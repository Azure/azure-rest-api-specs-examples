
/**
 * Samples for KubernetesVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/KubernetesVersions_Delete.json
     */
    /**
     * Sample code: Delete Kubernetes versions.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deleteKubernetesVersions(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.kubernetesVersions().delete("resourceGroupName", "default", null, null,
            com.azure.core.util.Context.NONE);
    }
}
