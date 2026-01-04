
/**
 * Samples for ManagedClusters ListKubernetesVersions.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * KubernetesVersions_List.json
     */
    /**
     * Sample code: List Kubernetes Versions.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listKubernetesVersions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getManagedClusters()
            .listKubernetesVersionsWithResponse("location1", com.azure.core.util.Context.NONE);
    }
}
