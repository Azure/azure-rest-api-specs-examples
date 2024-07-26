
/**
 * Samples for VirtualMachines ConvertToManagedDisks.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * virtualMachineExamples/VirtualMachine_ConvertToManagedDisks_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachine_ConvertToManagedDisks_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualMachineConvertToManagedDisksMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachines().convertToManagedDisks("rgcompute",
            "aaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
