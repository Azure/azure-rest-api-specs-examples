
/**
 * Samples for AgentPools Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/AgentPools_Get.json
     */
    /**
     * Sample code: Get Kubernetes cluster agent pool.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        getKubernetesClusterAgentPool(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.agentPools().getWithResponse("resourceGroupName", "kubernetesClusterName", "agentPoolName",
            com.azure.core.util.Context.NONE);
    }
}
