
/**
 * Samples for ManagedClusters RotateServiceAccountSigningKeys.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * ManagedClustersRotateServiceAccountSigningKeys.json
     */
    /**
     * Sample code: Rotate Cluster Service Account Signing Keys.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void rotateClusterServiceAccountSigningKeys(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getManagedClusters().rotateServiceAccountSigningKeys("rg1",
            "clustername1", com.azure.core.util.Context.NONE);
    }
}
