
/**
 * Samples for VirtualMachineImages ListSkus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineImageExamples/VirtualMachineImage_ListSkus_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImage_ListSkus_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineImageListSkusMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineImages().listSkusWithResponse("aaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaa",
            "aaaaaaa", com.azure.core.util.Context.NONE);
    }
}
