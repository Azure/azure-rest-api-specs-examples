import com.azure.core.util.Context;

/** Samples for ElasticSans Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2021-11-20-preview/examples/ElasticSans_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: ElasticSans_Delete_MaximumSet_Gen.
     *
     * @param manager Entry point to ElasticSanManager.
     */
    public static void elasticSansDeleteMaximumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.elasticSans().delete("rgelasticsan", "ti7q-k952-1qB3J_5", Context.NONE);
    }
}
