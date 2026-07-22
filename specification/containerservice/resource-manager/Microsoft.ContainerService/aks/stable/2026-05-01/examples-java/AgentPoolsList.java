
/**
 * Samples for AgentPools List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/AgentPoolsList.json
     */
    /**
     * Sample code: List Agent Pools by Managed Cluster.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        listAgentPoolsByManagedCluster(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getAgentPools().list("rg1", "clustername1", com.azure.core.util.Context.NONE);
    }
}
