/** Samples for AgentVersion List. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/AgentVersions_Get.json
     */
    /**
     * Sample code: GET Agent Versions.
     *
     * @param manager Entry point to HybridComputeManager.
     */
    public static void gETAgentVersions(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.agentVersions().listWithResponse("myOsType", com.azure.core.util.Context.NONE);
    }
}
