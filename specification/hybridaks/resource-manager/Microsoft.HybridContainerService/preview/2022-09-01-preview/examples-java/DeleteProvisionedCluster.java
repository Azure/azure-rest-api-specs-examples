/** Samples for ProvisionedClustersOperation Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-09-01-preview/examples/DeleteProvisionedCluster.json
     */
    /**
     * Sample code: DeleteProvisionedCluster.
     *
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void deleteProvisionedCluster(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager
            .provisionedClustersOperations()
            .deleteByResourceGroupWithResponse(
                "test-arcappliance-resgrp", "test-hybridakscluster", com.azure.core.util.Context.NONE);
    }
}
