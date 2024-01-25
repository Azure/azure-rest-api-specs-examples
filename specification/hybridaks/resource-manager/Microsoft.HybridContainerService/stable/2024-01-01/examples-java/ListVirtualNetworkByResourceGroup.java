
/**
 * Samples for VirtualNetworks ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/
     * ListVirtualNetworkByResourceGroup.json
     */
    /**
     * Sample code: ListVirtualNetworkByResourceGroup.
     * 
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void listVirtualNetworkByResourceGroup(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager.virtualNetworks().listByResourceGroup("test-arcappliance-resgrp", com.azure.core.util.Context.NONE);
    }
}
