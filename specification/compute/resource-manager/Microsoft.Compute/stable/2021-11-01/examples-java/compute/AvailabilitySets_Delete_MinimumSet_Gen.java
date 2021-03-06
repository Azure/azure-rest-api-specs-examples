import com.azure.core.util.Context;

/** Samples for AvailabilitySets Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/AvailabilitySets_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: AvailabilitySets_Delete_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void availabilitySetsDeleteMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getAvailabilitySets()
            .deleteWithResponse("rgcompute", "aaaaaaaaaaa", Context.NONE);
    }
}
