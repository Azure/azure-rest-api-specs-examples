
/**
 * Samples for AvailabilitySets ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * availabilitySetExamples/AvailabilitySet_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: AvailabilitySet_List_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void availabilitySetListMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getAvailabilitySets().listByResourceGroup("rgcompute",
            com.azure.core.util.Context.NONE);
    }
}
