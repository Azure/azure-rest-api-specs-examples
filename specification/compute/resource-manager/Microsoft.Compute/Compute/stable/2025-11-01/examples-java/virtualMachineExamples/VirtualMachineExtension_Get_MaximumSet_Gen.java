
/**
 * Samples for VirtualMachineExtensions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineExamples/VirtualMachineExtension_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineExtension_Get_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineExtensionGetMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineExtensions().getWithResponse("rgcompute",
            "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaa", "aaaaaa", com.azure.core.util.Context.NONE);
    }
}
