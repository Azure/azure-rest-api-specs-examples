/** Samples for ResourceSkus List. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/Skus/stable/2021-07-01/examples/skus/ListAvailableResourceSkusForARegion.json
     */
    /**
     * Sample code: Lists all available Resource SKUs for the specified region.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listsAllAvailableResourceSKUsForTheSpecifiedRegion(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getResourceSkus()
            .list("location eq 'westus'", null, com.azure.core.util.Context.NONE);
    }
}
