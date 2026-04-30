
/**
 * Samples for VirtualMachineImages ListPublishers.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineImageExamples/VirtualMachineImage_ListPublishers_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImage_ListPublishers_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineImageListPublishersMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineImages().listPublishersWithResponse("aaaaaaaaaaaaaaaaaaaaaaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
