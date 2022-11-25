import com.azure.core.util.Context;

/** Samples for VMCollection Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/elastic/resource-manager/Microsoft.Elastic/preview/2022-07-01-preview/examples/VMCollection_Update.json
     */
    /**
     * Sample code: VMCollection_Update.
     *
     * @param manager Entry point to ElasticManager.
     */
    public static void vMCollectionUpdate(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.vMCollections().updateWithResponse("myResourceGroup", "myMonitor", null, Context.NONE);
    }
}
