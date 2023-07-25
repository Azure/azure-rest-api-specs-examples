/** Samples for Skus List. */
public final class Main {
    /*
     * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2022-12-01-preview/examples/Skus_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: Skus_List_MinimumSet_Gen.
     *
     * @param manager Entry point to ElasticSanManager.
     */
    public static void skusListMinimumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.skus().list(null, com.azure.core.util.Context.NONE);
    }
}
