import com.azure.core.util.Context;

/** Samples for VirtualMachineImagesEdgeZone Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineImagesEdgeZone_Get_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineImagesEdgeZoneGetMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineImagesEdgeZones()
            .getWithResponse(
                "aaaaaaaaaaaaaaaaaaaaaaaa",
                "aaaaaaaa",
                "aaaaaaaaaaaaaaaaaaaaaaa",
                "aaaaaaaaaaaaaaaaaaaaaaaaaaa",
                "aaaaaaaaaa",
                "aaaaaaaaaaaaaaaaaaaaaaaaaaa",
                Context.NONE);
    }
}
