/** Samples for ProvisionedClustersOperation UpgradeNodeImageVersionForEntireCluster. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-09-01-preview/examples/ProvisionedClustersUpgradeNodeImageVersionForEntireCluster.json
     */
    /**
     * Sample code: UpgradeClusterNodeImageVersion.
     *
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void upgradeClusterNodeImageVersion(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager
            .provisionedClustersOperations()
            .upgradeNodeImageVersionForEntireCluster(
                "test-arcappliance-resgrp", "test-hybridakscluster", com.azure.core.util.Context.NONE);
    }
}
