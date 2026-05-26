
/**
 * Samples for MeshMemberships ListByManagedCluster.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/MeshMemberships_ListByManagedCluster.json
     */
    /**
     * Sample code: List Mesh Memberships by Managed Cluster.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void listMeshMembershipsByManagedCluster(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getMeshMemberships().listByManagedCluster("rg1", "clustername1",
            com.azure.core.util.Context.NONE);
    }
}
