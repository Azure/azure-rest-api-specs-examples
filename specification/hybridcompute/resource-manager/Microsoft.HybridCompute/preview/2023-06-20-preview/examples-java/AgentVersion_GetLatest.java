/** Samples for AgentVersion Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/AgentVersion_GetLatest.json
     */
    /**
     * Sample code: GET Agent Versions.
     *
     * @param manager Entry point to HybridComputeManager.
     */
    public static void gETAgentVersions(com.azure.resourcemanager.hybridcompute.HybridComputeManager manager) {
        manager.agentVersions().getWithResponse("myOsType", "1.27", com.azure.core.util.Context.NONE);
    }
}
