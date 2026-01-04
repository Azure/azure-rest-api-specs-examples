
/**
 * Samples for AgentPools Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/RegistryTasks/preview/2019-06-01-
     * preview/examples/AgentPoolsDelete.json
     */
    /**
     * Sample code: AgentPools_Delete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void agentPoolsDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getAgentPools().delete("myResourceGroup", "myRegistry",
            "myAgentPool", com.azure.core.util.Context.NONE);
    }
}
