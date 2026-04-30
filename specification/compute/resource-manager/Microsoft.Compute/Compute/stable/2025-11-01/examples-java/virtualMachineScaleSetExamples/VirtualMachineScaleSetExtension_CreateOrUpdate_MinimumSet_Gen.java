
import com.azure.resourcemanager.compute.fluent.models.VirtualMachineScaleSetExtensionInner;

/**
 * Samples for VirtualMachineScaleSetExtensions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetExtension_CreateOrUpdate_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetExtension_CreateOrUpdate_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineScaleSetExtensionCreateOrUpdateMinimumSetGen(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetExtensions().createOrUpdate("rgcompute", "aaaaaaaaaaa",
            "aaaaaaaaaaa", new VirtualMachineScaleSetExtensionInner(), com.azure.core.util.Context.NONE);
    }
}
