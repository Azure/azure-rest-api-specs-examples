import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2021-11-20-preview/examples/Operations_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Operations_List_MaximumSet_Gen.
     *
     * @param manager Entry point to ElasticSanManager.
     */
    public static void operationsListMaximumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.operations().list(Context.NONE);
    }
}
