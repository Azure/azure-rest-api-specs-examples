
/**
 * Samples for AgentPools Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * AgentPoolsGet.json
     */
    /**
     * Sample code: Get Agent Pool.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAgentPool(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getAgentPools().getWithResponse("rg1", "clustername1",
            "agentpool1", com.azure.core.util.Context.NONE);
    }
}
