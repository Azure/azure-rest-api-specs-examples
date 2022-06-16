import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetVMs Restart. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSetVMs_Restart_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVMs_Restart_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetVMsRestartMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetVMs()
            .restart("rgcompute", "aaaaaaaaaaaa", "aaaaaa", Context.NONE);
    }
}
