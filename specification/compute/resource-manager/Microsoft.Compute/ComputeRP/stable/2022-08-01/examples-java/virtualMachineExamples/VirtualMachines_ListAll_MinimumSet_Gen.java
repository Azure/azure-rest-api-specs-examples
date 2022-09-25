import com.azure.core.util.Context;

/** Samples for VirtualMachines List. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineExamples/VirtualMachines_ListAll_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachines_ListAll_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachinesListAllMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachines().list(null, null, Context.NONE);
    }
}
