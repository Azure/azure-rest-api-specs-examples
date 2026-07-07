
/**
 * Samples for Skus List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/SKUListWithLocationInfo.json
     */
    /**
     * Sample code: SKUListWithLocationInfo.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void sKUListWithLocationInfo(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getSkus().list(com.azure.core.util.Context.NONE);
    }
}
