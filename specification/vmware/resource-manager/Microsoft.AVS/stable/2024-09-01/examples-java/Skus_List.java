
/**
 * Samples for Skus List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/Skus_List.json
     */
    /**
     * Sample code: Skus_List.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void skusList(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.skus().list(com.azure.core.util.Context.NONE);
    }
}
