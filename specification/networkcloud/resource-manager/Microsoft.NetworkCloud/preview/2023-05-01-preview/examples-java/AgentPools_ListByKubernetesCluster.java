/** Samples for AgentPools ListByKubernetesCluster. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/AgentPools_ListByKubernetesCluster.json
     */
    /**
     * Sample code: List agent pools of the Kubernetes cluster.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void listAgentPoolsOfTheKubernetesCluster(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager
            .agentPools()
            .listByKubernetesCluster("resourceGroupName", "kubernetesClusterName", com.azure.core.util.Context.NONE);
    }
}
