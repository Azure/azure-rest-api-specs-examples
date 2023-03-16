/** Samples for ProvisionedClustersOperation GetUpgradeProfile. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-09-01-preview/examples/ProvisionedClustersGetUpgradeProfile.json
     */
    /**
     * Sample code: GetUpgradeProfileForProvisionedCluster.
     *
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void getUpgradeProfileForProvisionedCluster(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager
            .provisionedClustersOperations()
            .getUpgradeProfileWithResponse(
                "test-arcappliance-resgrp", "test-hybridakscluster", com.azure.core.util.Context.NONE);
    }
}
