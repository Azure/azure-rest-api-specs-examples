import com.azure.core.util.Context;

/** Samples for VirtualMachines ConvertToManagedDisks. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/virtualMachineExamples/VirtualMachines_ConvertToManagedDisks_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachines_ConvertToManagedDisks_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachinesConvertToManagedDisksMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachines()
            .convertToManagedDisks("rgcompute", "aaaaaaa", Context.NONE);
    }
}
