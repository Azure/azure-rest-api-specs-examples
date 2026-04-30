
/**
 * Samples for VirtualMachineScaleSets Reapply.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Reapply_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSets_Reapply_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetsReapplyMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().reapply("VirtualMachineScaleSetReapplyTestRG",
            "VMSSReapply-Test-ScaleSet", com.azure.core.util.Context.NONE);
    }
}
