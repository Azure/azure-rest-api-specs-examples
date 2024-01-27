
/**
 * Samples for VirtualMachineScaleSetVMs List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVM_List_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetVMListMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSetVMs().list("rgcompute",
            "aaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
