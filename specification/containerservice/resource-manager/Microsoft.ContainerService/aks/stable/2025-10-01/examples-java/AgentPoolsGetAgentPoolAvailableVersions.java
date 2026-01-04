
/**
 * Samples for AgentPools GetAvailableAgentPoolVersions.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * AgentPoolsGetAgentPoolAvailableVersions.json
     */
    /**
     * Sample code: Get available versions for agent pool.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAvailableVersionsForAgentPool(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getAgentPools()
            .getAvailableAgentPoolVersionsWithResponse("rg1", "clustername1", com.azure.core.util.Context.NONE);
    }
}
