
/**
 * Samples for NetworkInterfaces ListVirtualMachineScaleSetNetworkInterfaces.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/VmssNetworkInterfaceList.json
     */
    /**
     * Sample code: List virtual machine scale set network interfaces.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listVirtualMachineScaleSetNetworkInterfaces(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkInterfaces()
            .listVirtualMachineScaleSetNetworkInterfaces("rg1", "vmss1", com.azure.core.util.Context.NONE);
    }
}
