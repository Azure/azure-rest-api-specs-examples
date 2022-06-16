import com.azure.core.util.Context;

/** Samples for AgentPools GetUpgradeProfile. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-02-01/examples/AgentPoolsGetUpgradeProfile.json
     */
    /**
     * Sample code: Get Upgrade Profile for Agent Pool.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getUpgradeProfileForAgentPool(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .kubernetesClusters()
            .manager()
            .serviceClient()
            .getAgentPools()
            .getUpgradeProfileWithResponse("rg1", "clustername1", "agentpool1", Context.NONE);
    }
}
