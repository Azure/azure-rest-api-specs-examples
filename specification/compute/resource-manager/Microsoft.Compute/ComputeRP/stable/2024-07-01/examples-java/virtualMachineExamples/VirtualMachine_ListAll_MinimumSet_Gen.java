
/**
 * Samples for VirtualMachines List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * virtualMachineExamples/VirtualMachine_ListAll_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachine_ListAll_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineListAllMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachines().list(null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
