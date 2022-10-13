import com.azure.core.util.Context;
import com.azure.resourcemanager.hybridcontainerservice.models.AgentPool;

/** Samples for AgentPool Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-05-01-preview/examples/UpdateAgentPool.json
     */
    /**
     * Sample code: UpdateAgentPool.
     *
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void updateAgentPool(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        AgentPool resource =
            manager
                .agentPools()
                .getWithResponse(
                    "test-arcappliance-resgrp", "test-hybridakscluster", "test-hybridaksnodepool", Context.NONE)
                .getValue();
        resource.update().withCount(3).apply();
    }
}
