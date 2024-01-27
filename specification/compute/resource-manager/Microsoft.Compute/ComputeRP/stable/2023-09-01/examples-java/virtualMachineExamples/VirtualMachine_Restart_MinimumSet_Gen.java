
/**
 * Samples for VirtualMachines Restart.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/
     * virtualMachineExamples/VirtualMachine_Restart_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachine_Restart_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineRestartMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachines().restart("rgcompute", "aaa",
            com.azure.core.util.Context.NONE);
    }
}
