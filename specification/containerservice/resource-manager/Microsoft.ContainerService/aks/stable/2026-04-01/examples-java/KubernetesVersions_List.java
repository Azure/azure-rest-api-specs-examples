
/**
 * Samples for ManagedClusters ListKubernetesVersions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/KubernetesVersions_List.json
     */
    /**
     * Sample code: List Kubernetes Versions.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        listKubernetesVersions(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedClusters().listKubernetesVersionsWithResponse("location1",
            com.azure.core.util.Context.NONE);
    }
}
