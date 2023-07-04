/** Samples for AgentPools Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/AgentPools_Delete.json
     */
    /**
     * Sample code: Delete Kubernetes cluster agent pool.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deleteKubernetesClusterAgentPool(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager
            .agentPools()
            .delete("resourceGroupName", "kubernetesClusterName", "agentPoolName", com.azure.core.util.Context.NONE);
    }
}
