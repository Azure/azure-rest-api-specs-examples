
/**
 * Samples for NetworkInterfaces ListVirtualMachineScaleSetVMNetworkInterfaces.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VmssVmNetworkInterfaceList.
     * json
     */
    /**
     * Sample code: List virtual machine scale set vm network interfaces.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listVirtualMachineScaleSetVmNetworkInterfaces(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkInterfaces()
            .listVirtualMachineScaleSetVMNetworkInterfaces("rg1", "vmss1", "1", com.azure.core.util.Context.NONE);
    }
}
