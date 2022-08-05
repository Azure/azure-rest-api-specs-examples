import com.azure.core.util.Context;

/** Samples for VirtualMachines PowerOff. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineExamples/VirtualMachines_PowerOff_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachines_PowerOff_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachinesPowerOffMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachines()
            .powerOff("rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaaaaa", true, Context.NONE);
    }
}
