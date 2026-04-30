
/**
 * Samples for ResourceSkus List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2021-07-01/skus/ListAvailableResourceSkus.json
     */
    /**
     * Sample code: Lists all available Resource SKUs.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void listsAllAvailableResourceSKUs(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getResourceSkus().list(null, null, com.azure.core.util.Context.NONE);
    }
}
