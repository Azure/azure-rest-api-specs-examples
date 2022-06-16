import com.azure.core.util.Context;

/** Samples for AvailabilitySets GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/availabilitySetExamples/AvailabilitySets_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: AvailabilitySets_Get_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void availabilitySetsGetMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getAvailabilitySets()
            .getByResourceGroupWithResponse("rgcompute", "aaaaaaaaaaaaaaaaaaaa", Context.NONE);
    }
}
