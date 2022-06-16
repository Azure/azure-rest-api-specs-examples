import com.azure.core.util.Context;

/** Samples for VirtualMachineImagesEdgeZone List. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImagesEdgeZone_List_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineImagesEdgeZoneListMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineImagesEdgeZones()
            .listWithResponse(
                "aaaaaaaaaaaaaaa",
                "aaaaaaaaaaaaaaaaaaaaaaaaa",
                "aaaa",
                "aaaaaaaaaaaaaaaaaaaaaaaaaa",
                "aaaaaaaaaaaaaaaaaaaaaaa",
                "aaaaaaaaaaaaaaaaaaaaaaaa",
                12,
                "aaaaaaaaaaaaaaaaaaaaaaaaaaaa",
                Context.NONE);
    }
}
