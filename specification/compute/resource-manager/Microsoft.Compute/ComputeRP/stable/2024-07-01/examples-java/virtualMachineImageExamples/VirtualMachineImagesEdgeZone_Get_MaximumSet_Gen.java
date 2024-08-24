
/**
 * Samples for VirtualMachineImagesEdgeZone Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * virtualMachineImageExamples/VirtualMachineImagesEdgeZone_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImagesEdgeZone_Get_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualMachineImagesEdgeZoneGetMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineImagesEdgeZones().getWithResponse(
            "aaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaa",
            "aaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
