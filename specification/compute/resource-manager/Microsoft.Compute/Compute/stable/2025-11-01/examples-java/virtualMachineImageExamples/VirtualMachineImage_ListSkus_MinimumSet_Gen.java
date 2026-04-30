
/**
 * Samples for VirtualMachineImages ListSkus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineImageExamples/VirtualMachineImage_ListSkus_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImage_ListSkus_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineImageListSkusMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineImages().listSkusWithResponse("aaaa", "aaaaaaaaaaaaa", "aaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
