
/**
 * Samples for Skus List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/
     * Skus_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Skus_List_MaximumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void skusListMaximumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.skus().list("obwwdrkq", com.azure.core.util.Context.NONE);
    }
}
