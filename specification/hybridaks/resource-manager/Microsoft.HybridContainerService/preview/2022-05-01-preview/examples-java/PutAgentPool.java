import com.azure.resourcemanager.hybridcontainerservice.models.OsType;

/** Samples for AgentPool CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-05-01-preview/examples/PutAgentPool.json
     */
    /**
     * Sample code: PutAgentPool.
     *
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void putAgentPool(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager
            .agentPools()
            .define("test-hybridaksnodepool")
            .withRegion("westus")
            .withExistingProvisionedCluster("test-arcappliance-resgrp", "test-hybridakscluster")
            .withCount(1)
            .withOsType(OsType.LINUX)
            .withVmSize("Standard_A4_v2")
            .create();
    }
}
