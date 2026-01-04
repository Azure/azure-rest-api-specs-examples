
/**
 * Samples for NetworkInterfaces GetVirtualMachineScaleSetIpConfiguration.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * VmssNetworkInterfaceIpConfigGet.json
     */
    /**
     * Sample code: Get virtual machine scale set network interface.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVirtualMachineScaleSetNetworkInterface(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkInterfaces()
            .getVirtualMachineScaleSetIpConfigurationWithResponse("rg1", "vmss1", "2", "nic1", "ip1", null,
                com.azure.core.util.Context.NONE);
    }
}
