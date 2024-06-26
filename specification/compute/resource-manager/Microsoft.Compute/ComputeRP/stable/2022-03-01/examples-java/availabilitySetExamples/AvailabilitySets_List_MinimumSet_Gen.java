import com.azure.core.util.Context;

/** Samples for AvailabilitySets ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/availabilitySetExamples/AvailabilitySets_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: AvailabilitySets_List_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void availabilitySetsListMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getAvailabilitySets()
            .listByResourceGroup("rgcompute", Context.NONE);
    }
}
