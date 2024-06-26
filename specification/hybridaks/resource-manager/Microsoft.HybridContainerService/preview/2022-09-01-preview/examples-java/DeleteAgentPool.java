/** Samples for AgentPool Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-09-01-preview/examples/DeleteAgentPool.json
     */
    /**
     * Sample code: DeleteAgentPool.
     *
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void deleteAgentPool(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager
            .agentPools()
            .deleteWithResponse(
                "test-arcappliance-resgrp",
                "test-hybridakscluster",
                "test-hybridaksnodepool",
                com.azure.core.util.Context.NONE);
    }
}
