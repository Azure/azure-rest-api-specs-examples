import com.azure.core.util.Context;

/** Samples for NetworkInterfaces ListVirtualMachineScaleSetIpConfigurations. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VmssNetworkInterfaceIpConfigList.json
     */
    /**
     * Sample code: List virtual machine scale set network interface ip configurations.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listVirtualMachineScaleSetNetworkInterfaceIpConfigurations(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkInterfaces()
            .listVirtualMachineScaleSetIpConfigurations("rg1", "vmss1", "2", "nic1", null, Context.NONE);
    }
}
