import com.azure.core.util.Context;
import com.azure.resourcemanager.elasticsan.models.ElasticSan;
import java.util.HashMap;
import java.util.Map;

/** Samples for ElasticSans Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2021-11-20-preview/examples/ElasticSans_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: ElasticSans_Update_MaximumSet_Gen.
     *
     * @param manager Entry point to ElasticSanManager.
     */
    public static void elasticSansUpdateMaximumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        ElasticSan resource =
            manager
                .elasticSans()
                .getByResourceGroupWithResponse("rgelasticsan", "ti7q-k952-1qB3J_5", Context.NONE)
                .getValue();
        resource.update().withTags(mapOf("key3137", "aaaaaaaaaaaaaaa")).apply();
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
