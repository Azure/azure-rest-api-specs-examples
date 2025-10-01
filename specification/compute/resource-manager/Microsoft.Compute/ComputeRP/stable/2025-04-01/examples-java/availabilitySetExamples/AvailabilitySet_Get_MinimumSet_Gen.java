
/**
 * Samples for AvailabilitySets GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/
     * availabilitySetExamples/AvailabilitySet_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: AvailabilitySet_Get_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void availabilitySetGetMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getAvailabilitySets()
            .getByResourceGroupWithResponse("rgcompute", "aaaaaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
