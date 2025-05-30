
/**
 * Samples for VirtualMachineScaleSets Reapply.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSet_Reapply_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSets_Reapply_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualMachineScaleSetsReapplyMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSets().reapply(
            "VirtualMachineScaleSetReapplyTestRG", "VMSSReapply-Test-ScaleSet", com.azure.core.util.Context.NONE);
    }
}
