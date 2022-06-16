import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSets Restart. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSets_Restart_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSets_Restart_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetsRestartMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSets()
            .restart("rgcompute", "aaaa", null, Context.NONE);
    }
}
