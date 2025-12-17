
/**
 * Samples for AgentPools Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/AgentPools_Delete.
     * json
     */
    /**
     * Sample code: Delete Kubernetes cluster agent pool.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        deleteKubernetesClusterAgentPool(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.agentPools().delete("resourceGroupName", "kubernetesClusterName", "agentPoolName", null, null,
            com.azure.core.util.Context.NONE);
    }
}
