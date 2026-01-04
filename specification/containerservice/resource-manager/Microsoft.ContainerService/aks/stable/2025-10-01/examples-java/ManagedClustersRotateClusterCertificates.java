
/**
 * Samples for ManagedClusters RotateClusterCertificates.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * ManagedClustersRotateClusterCertificates.json
     */
    /**
     * Sample code: Rotate Cluster Certificates.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void rotateClusterCertificates(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getManagedClusters().rotateClusterCertificates("rg1",
            "clustername1", com.azure.core.util.Context.NONE);
    }
}
