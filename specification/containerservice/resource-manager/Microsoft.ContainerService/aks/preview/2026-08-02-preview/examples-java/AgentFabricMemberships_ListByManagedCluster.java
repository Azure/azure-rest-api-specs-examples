
/**
 * Samples for AgentFabricMemberships ListByManagedCluster.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-02-preview/AgentFabricMemberships_ListByManagedCluster.json
     */
    /**
     * Sample code: List AgentFabric Memberships by Managed Cluster.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void listAgentFabricMembershipsByManagedCluster(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getAgentFabricMemberships().listByManagedCluster("rg1", "clustername1",
            com.azure.core.util.Context.NONE);
    }
}
