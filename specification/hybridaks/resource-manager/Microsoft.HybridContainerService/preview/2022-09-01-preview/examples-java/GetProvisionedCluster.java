/** Samples for ProvisionedClustersOperation GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-09-01-preview/examples/GetProvisionedCluster.json
     */
    /**
     * Sample code: GetProvisionedCluster.
     *
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void getProvisionedCluster(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager
            .provisionedClustersOperations()
            .getByResourceGroupWithResponse(
                "test-arcappliance-resgrp", "test-hybridakscluster", com.azure.core.util.Context.NONE);
    }
}
