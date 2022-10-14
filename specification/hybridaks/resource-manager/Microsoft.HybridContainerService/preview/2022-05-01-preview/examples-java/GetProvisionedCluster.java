import com.azure.core.util.Context;

/** Samples for ProvisionedClustersOperation GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-05-01-preview/examples/GetProvisionedCluster.json
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
            .getByResourceGroupWithResponse("test-arcappliance-resgrp", "test-hybridakscluster", Context.NONE);
    }
}
