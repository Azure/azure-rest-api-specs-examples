
/**
 * Samples for KubernetesVersions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/KubernetesVersions_ListBySubscription.json
     */
    /**
     * Sample code: List Kubernetes versions for subscription.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        listKubernetesVersionsForSubscription(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.kubernetesVersions().list(null, null, com.azure.core.util.Context.NONE);
    }
}
