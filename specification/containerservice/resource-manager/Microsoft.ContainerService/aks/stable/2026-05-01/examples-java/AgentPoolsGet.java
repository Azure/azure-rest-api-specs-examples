
/**
 * Samples for AgentPools Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/AgentPoolsGet.json
     */
    /**
     * Sample code: Get Agent Pool.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void getAgentPool(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getAgentPools().getWithResponse("rg1", "clustername1", "agentpool1",
            com.azure.core.util.Context.NONE);
    }
}
