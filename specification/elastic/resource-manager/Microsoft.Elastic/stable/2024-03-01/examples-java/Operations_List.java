
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/elastic/resource-manager/Microsoft.Elastic/stable/2024-03-01/examples/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     * 
     * @param manager Entry point to ElasticManager.
     */
    public static void operationsList(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
