import com.azure.core.util.Context;

/** Samples for PublicIpAddresses GetVirtualMachineScaleSetPublicIpAddress. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/VmssPublicIpGet.json
     */
    /**
     * Sample code: GetVMSSPublicIP.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVMSSPublicIP(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getPublicIpAddresses()
            .getVirtualMachineScaleSetPublicIpAddressWithResponse(
                "vmss-tester", "vmss1", "1", "nic1", "ip1", "pub1", null, Context.NONE);
    }
}
