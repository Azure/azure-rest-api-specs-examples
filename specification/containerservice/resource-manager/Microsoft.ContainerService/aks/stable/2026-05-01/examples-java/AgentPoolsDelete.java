
/**
 * Samples for AgentPools Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/AgentPoolsDelete.json
     */
    /**
     * Sample code: Delete Agent Pool.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void deleteAgentPool(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getAgentPools().delete("rg1", "clustername1", "agentpool1", null, null,
            com.azure.core.util.Context.NONE);
    }
}
