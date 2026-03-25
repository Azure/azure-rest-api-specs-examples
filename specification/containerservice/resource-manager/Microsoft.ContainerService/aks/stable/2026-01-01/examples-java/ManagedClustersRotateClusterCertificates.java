
/**
 * Samples for ManagedClusters RotateClusterCertificates.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/ManagedClustersRotateClusterCertificates.json
     */
    /**
     * Sample code: Rotate Cluster Certificates.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        rotateClusterCertificates(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedClusters().rotateClusterCertificates("rg1", "clustername1",
            com.azure.core.util.Context.NONE);
    }
}
