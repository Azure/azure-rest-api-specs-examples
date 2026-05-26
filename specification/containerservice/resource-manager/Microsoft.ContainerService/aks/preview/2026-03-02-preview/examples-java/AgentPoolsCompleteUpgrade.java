
/**
 * Samples for AgentPools CompleteUpgrade.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/AgentPoolsCompleteUpgrade.json
     */
    /**
     * Sample code: Complete agent pool upgrade.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        completeAgentPoolUpgrade(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getAgentPools().completeUpgrade("rg1", "clustername1", "agentpool1",
            com.azure.core.util.Context.NONE);
    }
}
