/** Samples for VirtualMachineScaleSetVMs Restart. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Restart_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVM_Restart_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetVMRestartMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetVMs()
            .restart("rgcompute", "aaaaaaaaaaaa", "aaaaaa", com.azure.core.util.Context.NONE);
    }
}
