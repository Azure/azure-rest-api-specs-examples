
/**
 * Samples for AgentPool Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/
     * DeleteAgentPool.json
     */
    /**
     * Sample code: DeleteAgentPool.
     * 
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void
        deleteAgentPool(com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager.agentPools().delete(
            "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/test-hybridakscluster",
            "testnodepool", com.azure.core.util.Context.NONE);
    }
}
