
/**
 * Samples for ResourceSkus List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2021-07-01/skus/ListAvailableResourceSkusWithExtendedLocations.json
     */
    /**
     * Sample code: Lists all available Resource SKUs with Extended Location information.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void listsAllAvailableResourceSKUsWithExtendedLocationInformation(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getResourceSkus().list(null, "true", com.azure.core.util.Context.NONE);
    }
}
