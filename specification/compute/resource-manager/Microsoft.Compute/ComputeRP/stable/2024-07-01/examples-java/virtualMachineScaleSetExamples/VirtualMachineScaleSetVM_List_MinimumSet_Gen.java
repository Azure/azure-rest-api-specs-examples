
/**
 * Samples for VirtualMachineScaleSetVMs List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVM_List_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetVMListMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSetVMs().list("rgcompute",
            "aaaaaaaaaaaaaaaaaaaaaaaaaaaa", null, null, null, com.azure.core.util.Context.NONE);
    }
}
