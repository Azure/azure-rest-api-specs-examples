
/**
 * Samples for ManagedClusters RotateServiceAccountSigningKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/ManagedClustersRotateServiceAccountSigningKeys.json
     */
    /**
     * Sample code: Rotate Cluster Service Account Signing Keys.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void rotateClusterServiceAccountSigningKeys(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedClusters().rotateServiceAccountSigningKeys("rg1", "clustername1",
            com.azure.core.util.Context.NONE);
    }
}
