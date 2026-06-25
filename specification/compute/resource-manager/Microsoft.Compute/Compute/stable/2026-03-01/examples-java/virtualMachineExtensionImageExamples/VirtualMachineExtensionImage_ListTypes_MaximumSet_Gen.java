
/**
 * Samples for VirtualMachineExtensionImages ListTypes.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-03-01/virtualMachineExtensionImageExamples/VirtualMachineExtensionImage_ListTypes_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineExtensionImage_ListTypes_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineExtensionImageListTypesMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineExtensionImages().listTypesWithResponse("aaaaaaaaaaaaaaaaaaaaaaaaaa",
            "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
