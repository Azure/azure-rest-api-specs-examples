import com.azure.core.util.Context;

/** Samples for AvailabilitySets ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/availabilitySetExamples/AvailabilitySets_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: AvailabilitySets_List_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void availabilitySetsListMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getAvailabilitySets()
            .listByResourceGroup("rgcompute", Context.NONE);
    }
}
