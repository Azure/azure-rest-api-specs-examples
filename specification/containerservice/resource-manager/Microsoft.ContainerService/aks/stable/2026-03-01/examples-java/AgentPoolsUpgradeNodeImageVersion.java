
/**
 * Samples for AgentPools UpgradeNodeImageVersion.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/AgentPoolsUpgradeNodeImageVersion.json
     */
    /**
     * Sample code: Upgrade Agent Pool Node Image Version.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        upgradeAgentPoolNodeImageVersion(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getAgentPools().upgradeNodeImageVersion("rg1", "clustername1", "agentpool1",
            com.azure.core.util.Context.NONE);
    }
}
