
/**
 * Samples for VirtualMachineImages ListPublishers.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * virtualMachineImageExamples/VirtualMachineImage_ListPublishers_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImage_ListPublishers_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualMachineImageListPublishersMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineImages()
            .listPublishersWithResponse("aaaaaaaaaaaaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
