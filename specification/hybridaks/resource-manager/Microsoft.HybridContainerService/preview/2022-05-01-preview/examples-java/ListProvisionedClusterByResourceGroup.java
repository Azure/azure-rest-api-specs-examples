import com.azure.core.util.Context;

/** Samples for ProvisionedClustersOperation ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-05-01-preview/examples/ListProvisionedClusterByResourceGroup.json
     */
    /**
     * Sample code: ListProvisionedClusterByResourceGroup.
     *
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void listProvisionedClusterByResourceGroup(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager.provisionedClustersOperations().listByResourceGroup("test-arcappliance-resgrp", Context.NONE);
    }
}
