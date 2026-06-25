
/**
 * Samples for VirtualMachineImages ListPublishers.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineImageExamples/VirtualMachineImage_ListPublishers_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImage_ListPublishers_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineImageListPublishersMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineImages().listPublishersWithResponse("aaaaa",
            com.azure.core.util.Context.NONE);
    }
}
