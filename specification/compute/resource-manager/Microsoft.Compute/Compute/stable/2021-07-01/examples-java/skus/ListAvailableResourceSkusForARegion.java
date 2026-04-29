
/**
 * Samples for ResourceSkus List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2021-07-01/skus/ListAvailableResourceSkusForARegion.json
     */
    /**
     * Sample code: Lists all available Resource SKUs for the specified region.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        listsAllAvailableResourceSKUsForTheSpecifiedRegion(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getResourceSkus().list("location eq 'westus'", null, com.azure.core.util.Context.NONE);
    }
}
