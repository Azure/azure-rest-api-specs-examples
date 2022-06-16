import com.azure.core.util.Context;

/** Samples for VirtualMachines ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachines_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachines_List_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachinesListMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachines()
            .listByResourceGroup("rgcompute", "aaaaaaaaaaaaaaaaaaaaaaa", Context.NONE);
    }
}
