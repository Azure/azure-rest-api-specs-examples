
/**
 * Samples for AgentPools GetAvailableAgentPoolVersions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-02-preview/AgentPoolsGetAgentPoolAvailableVersions.json
     */
    /**
     * Sample code: Get available versions for agent pool.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        getAvailableVersionsForAgentPool(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getAgentPools().getAvailableAgentPoolVersionsWithResponse("rg1", "clustername1",
            com.azure.core.util.Context.NONE);
    }
}
