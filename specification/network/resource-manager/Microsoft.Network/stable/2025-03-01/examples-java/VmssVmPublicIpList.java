
/**
 * Samples for PublicIpAddresses ListVirtualMachineScaleSetVMPublicIpAddresses.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VmssVmPublicIpList.json
     */
    /**
     * Sample code: ListVMSSVMPublicIP.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listVMSSVMPublicIP(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPublicIpAddresses().listVirtualMachineScaleSetVMPublicIpAddresses(
            "vmss-tester", "vmss1", "1", "nic1", "ip1", com.azure.core.util.Context.NONE);
    }
}
