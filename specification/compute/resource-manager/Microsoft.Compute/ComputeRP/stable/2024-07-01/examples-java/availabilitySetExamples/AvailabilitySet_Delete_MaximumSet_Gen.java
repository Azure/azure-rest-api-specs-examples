
/**
 * Samples for AvailabilitySets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * availabilitySetExamples/AvailabilitySet_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: AvailabilitySet_Delete_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void availabilitySetDeleteMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getAvailabilitySets().deleteWithResponse("rgcompute",
            "aaaaaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
