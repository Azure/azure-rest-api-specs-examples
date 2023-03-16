/** Samples for AgentPool Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-09-01-preview/examples/GetAgentPool.json
     */
    /**
     * Sample code: GetAgentPool.
     *
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void getAgentPool(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager
            .agentPools()
            .getWithResponse(
                "test-arcappliance-resgrp",
                "test-hybridakscluster",
                "test-hybridaksnodepool",
                com.azure.core.util.Context.NONE);
    }
}
