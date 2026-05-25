
/**
 * Samples for AgentPools Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/AgentPoolsDelete_IgnorePodDisruptionBudget.json
     */
    /**
     * Sample code: Delete Agent Pool by ignoring PodDisruptionBudget.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void deleteAgentPoolByIgnoringPodDisruptionBudget(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getAgentPools().delete("rg1", "clustername1", "agentpool1", true, null,
            com.azure.core.util.Context.NONE);
    }
}
