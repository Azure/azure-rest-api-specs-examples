/** Samples for AvailabilitySets GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-11-01/examples/availabilitySetExamples/AvailabilitySets_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: AvailabilitySets_Get_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void availabilitySetsGetMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getAvailabilitySets()
            .getByResourceGroupWithResponse("rgcompute", "aaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
