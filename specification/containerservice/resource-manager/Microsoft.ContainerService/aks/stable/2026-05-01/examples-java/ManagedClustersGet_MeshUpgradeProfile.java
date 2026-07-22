
/**
 * Samples for ManagedClusters GetMeshUpgradeProfile.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/ManagedClustersGet_MeshUpgradeProfile.json
     */
    /**
     * Sample code: Gets version compatibility and upgrade profile for a service mesh in a cluster.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void getsVersionCompatibilityAndUpgradeProfileForAServiceMeshInACluster(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedClusters().getMeshUpgradeProfileWithResponse("rg1", "clustername1", "istio",
            com.azure.core.util.Context.NONE);
    }
}
