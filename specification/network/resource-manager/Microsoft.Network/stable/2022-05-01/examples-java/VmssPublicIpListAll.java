import com.azure.core.util.Context;

/** Samples for PublicIpAddresses ListVirtualMachineScaleSetPublicIpAddresses. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/VmssPublicIpListAll.json
     */
    /**
     * Sample code: ListVMSSPublicIP.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listVMSSPublicIP(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getPublicIpAddresses()
            .listVirtualMachineScaleSetPublicIpAddresses("vmss-tester", "vmss1", Context.NONE);
    }
}
