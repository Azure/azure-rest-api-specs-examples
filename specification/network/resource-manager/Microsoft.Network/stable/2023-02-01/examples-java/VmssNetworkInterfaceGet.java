/** Samples for NetworkInterfaces GetVirtualMachineScaleSetNetworkInterface. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/examples/VmssNetworkInterfaceGet.json
     */
    /**
     * Sample code: Get virtual machine scale set network interface.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVirtualMachineScaleSetNetworkInterface(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getNetworkInterfaces()
            .getVirtualMachineScaleSetNetworkInterfaceWithResponse(
                "rg1", "vmss1", "1", "nic1", null, com.azure.core.util.Context.NONE);
    }
}
