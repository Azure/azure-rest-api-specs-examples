
/**
 * Samples for VirtualMachines Start.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * virtualMachineExamples/VirtualMachine_Start_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachine_Start_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineStartMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachines().start("rgcompute", "aaaaa",
            com.azure.core.util.Context.NONE);
    }
}
