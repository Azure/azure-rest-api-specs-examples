
/**
 * Samples for VirtualNetworks Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/
     * DeleteVirtualNetwork.json
     */
    /**
     * Sample code: DeleteVirtualNetwork.
     * 
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void
        deleteVirtualNetwork(com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager.virtualNetworks().delete("test-arcappliance-resgrp", "test-vnet-static",
            com.azure.core.util.Context.NONE);
    }
}
