
/**
 * Samples for Location ListSupportedCloudServiceSkus.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/batch/resource-manager/Microsoft.Batch/stable/2024-02-01/examples/LocationListCloudServiceSkus.json
     */
    /**
     * Sample code: LocationListCloudServiceSkus.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void locationListCloudServiceSkus(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.locations().listSupportedCloudServiceSkus("japaneast", null, null, com.azure.core.util.Context.NONE);
    }
}
