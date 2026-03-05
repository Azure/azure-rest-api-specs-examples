
/**
 * Samples for ElasticVersions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/ElasticVersions_List.json
     */
    /**
     * Sample code: ElasticVersions_List.
     * 
     * @param manager Entry point to ElasticManager.
     */
    public static void elasticVersionsList(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.elasticVersions().list("myregion", com.azure.core.util.Context.NONE);
    }
}
