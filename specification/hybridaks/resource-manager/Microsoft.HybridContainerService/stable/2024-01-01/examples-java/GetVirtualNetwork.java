
/**
 * Samples for VirtualNetworks GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/
     * GetVirtualNetwork.json
     */
    /**
     * Sample code: GetVirtualNetwork.
     * 
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void
        getVirtualNetwork(com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager.virtualNetworks().getByResourceGroupWithResponse("test-arcappliance-resgrp", "test-vnet-static",
            com.azure.core.util.Context.NONE);
    }
}
