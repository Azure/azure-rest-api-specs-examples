
/**
 * Samples for AgentPools GetQueueStatus.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/RegistryTasks/preview/2019-06-01-
     * preview/examples/AgentPoolsGetQueueStatus.json
     */
    /**
     * Sample code: AgentPools_GetQueueStatus.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void agentPoolsGetQueueStatus(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getAgentPools().getQueueStatusWithResponse(
            "myResourceGroup", "myRegistry", "myAgentPool", com.azure.core.util.Context.NONE);
    }
}
