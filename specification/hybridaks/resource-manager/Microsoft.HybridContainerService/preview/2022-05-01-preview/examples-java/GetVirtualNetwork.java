import com.azure.core.util.Context;

/** Samples for VirtualNetworksOperation GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-05-01-preview/examples/GetVirtualNetwork.json
     */
    /**
     * Sample code: GetVirtualNetwork.
     *
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void getVirtualNetwork(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager
            .virtualNetworksOperations()
            .getByResourceGroupWithResponse("test-arcappliance-resgrp", "test-vnet-static", Context.NONE);
    }
}
