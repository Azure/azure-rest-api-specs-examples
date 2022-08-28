import com.azure.core.util.Context;

/** Samples for VirtualMachines PowerOff. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/virtualMachineExamples/VirtualMachines_PowerOff_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachines_PowerOff_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachinesPowerOffMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachines()
            .powerOff("rgcompute", "aaaaaaaaaaaaaaaaaa", null, Context.NONE);
    }
}
