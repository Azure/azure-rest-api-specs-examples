
import com.azure.resourcemanager.compute.models.VirtualMachineScaleSetExtensionUpdate;

/**
 * Samples for VirtualMachineScaleSetExtensions Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetExtension_Update_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetExtension_Update_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetExtensionUpdateMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetExtensions().update("rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaaaa",
            "aa", new VirtualMachineScaleSetExtensionUpdate(), com.azure.core.util.Context.NONE);
    }
}
