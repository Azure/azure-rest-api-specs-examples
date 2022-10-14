import com.azure.core.util.Context;

/** Samples for AgentPool ListByProvisionedCluster. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-05-01-preview/examples/ListAgentPoolByProvisionedCluster.json
     */
    /**
     * Sample code: ListAgentPoolByProvisionedCluster.
     *
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void listAgentPoolByProvisionedCluster(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager
            .agentPools()
            .listByProvisionedClusterWithResponse("test-arcappliance-resgrp", "test-hybridakscluster", Context.NONE);
    }
}
