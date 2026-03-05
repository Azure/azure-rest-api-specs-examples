
/**
 * Samples for VMCollection Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/VMCollection_Update.json
     */
    /**
     * Sample code: VMCollection_Update.
     * 
     * @param manager Entry point to ElasticManager.
     */
    public static void vMCollectionUpdate(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.vMCollections().updateWithResponse("myResourceGroup", "myMonitor", null,
            com.azure.core.util.Context.NONE);
    }
}
