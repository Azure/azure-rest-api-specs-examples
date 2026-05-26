
/**
 * Samples for AgentPools GetUpgradeProfile.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/AgentPoolsGetUpgradeProfile.json
     */
    /**
     * Sample code: Get Upgrade Profile for Agent Pool.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        getUpgradeProfileForAgentPool(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getAgentPools().getUpgradeProfileWithResponse("rg1", "clustername1", "agentpool1",
            com.azure.core.util.Context.NONE);
    }
}
