
/**
 * Samples for Skus List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/SKUList.json
     */
    /**
     * Sample code: SkuList.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void skuList(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getSkus().list(com.azure.core.util.Context.NONE);
    }
}
