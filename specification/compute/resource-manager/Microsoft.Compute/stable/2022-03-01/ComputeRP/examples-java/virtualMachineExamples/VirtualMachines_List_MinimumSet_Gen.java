import com.azure.core.util.Context;

/** Samples for VirtualMachines ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineExamples/VirtualMachines_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachines_List_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachinesListMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachines()
            .listByResourceGroup("rgcompute", null, Context.NONE);
    }
}
