
/**
 * Samples for AvailabilitySets GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * availabilitySetExamples/AvailabilitySet_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: AvailabilitySet_Get_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void availabilitySetGetMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getAvailabilitySets()
            .getByResourceGroupWithResponse("rgcompute", "aaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
