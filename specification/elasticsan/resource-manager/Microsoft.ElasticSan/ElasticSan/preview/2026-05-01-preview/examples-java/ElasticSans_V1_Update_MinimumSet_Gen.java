
import com.azure.resourcemanager.elasticsan.models.ElasticSan;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ElasticSans Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/ElasticSans_V1_Update_MinimumSet_Gen.json
     */
    /**
     * Sample code: ElasticSans_V1_Update_MinimumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void
        elasticSansV1UpdateMinimumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        ElasticSan resource = manager.elasticSans()
            .getByResourceGroupWithResponse("resourcegroupname", "elasticsanname", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().apply();
    }

    // Use "Map.of" if available
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
