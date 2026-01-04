
import com.azure.resourcemanager.containerregistry.models.AgentPoolUpdateParameters;

/**
 * Samples for AgentPools Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/RegistryTasks/preview/2019-06-01-
     * preview/examples/AgentPoolsUpdate.json
     */
    /**
     * Sample code: AgentPools_Update.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void agentPoolsUpdate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getAgentPools().update("myResourceGroup", "myRegistry",
            "myAgentPool", new AgentPoolUpdateParameters().withCount(1), com.azure.core.util.Context.NONE);
    }
}
