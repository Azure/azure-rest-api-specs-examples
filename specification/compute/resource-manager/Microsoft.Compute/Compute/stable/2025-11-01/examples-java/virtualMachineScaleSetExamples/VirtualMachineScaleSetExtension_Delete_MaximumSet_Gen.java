
/**
 * Samples for VirtualMachineScaleSetExtensions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetExtension_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetExtension_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetExtensionDeleteMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetExtensions().delete("rgcompute",
            "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
