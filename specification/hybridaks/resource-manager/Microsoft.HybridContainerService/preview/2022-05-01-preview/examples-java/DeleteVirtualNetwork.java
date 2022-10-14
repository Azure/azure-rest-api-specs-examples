import com.azure.core.util.Context;

/** Samples for VirtualNetworksOperation Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-05-01-preview/examples/DeleteVirtualNetwork.json
     */
    /**
     * Sample code: DeleteVirtualNetwork.
     *
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void deleteVirtualNetwork(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager
            .virtualNetworksOperations()
            .deleteByResourceGroupWithResponse("test-arcappliance-resgrp", "test-vnet-static", Context.NONE);
    }
}
