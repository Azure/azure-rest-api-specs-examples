
/**
 * Samples for AgentPools Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/RegistryTasks/preview/2019-06-01-
     * preview/examples/AgentPoolsGet.json
     */
    /**
     * Sample code: AgentPools_Get.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void agentPoolsGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getAgentPools().getWithResponse("myResourceGroup",
            "myRegistry", "myAgentPool", com.azure.core.util.Context.NONE);
    }
}
