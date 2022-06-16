import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetVMs RetrieveBootDiagnosticsData. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/compute/RetrieveBootDiagnosticsDataVMScaleSetVM.json
     */
    /**
     * Sample code: RetrieveBootDiagnosticsData of a virtual machine.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void retrieveBootDiagnosticsDataOfAVirtualMachine(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetVMs()
            .retrieveBootDiagnosticsDataWithResponse("ResourceGroup", "myvmScaleSet", "0", 60, Context.NONE);
    }
}
