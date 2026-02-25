
import com.azure.resourcemanager.elasticsan.models.AutoScalePolicyEnforcement;
import com.azure.resourcemanager.elasticsan.models.AutoScaleProperties;
import com.azure.resourcemanager.elasticsan.models.ElasticSan;
import com.azure.resourcemanager.elasticsan.models.PublicNetworkAccess;
import com.azure.resourcemanager.elasticsan.models.ScaleUpProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ElasticSans Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ElasticSans_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: ElasticSans_Update_MaximumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void elasticSansUpdateMaximumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        ElasticSan resource = manager.elasticSans()
            .getByResourceGroupWithResponse("resourcegroupname", "elasticsanname", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("key1931", "fakeTokenPlaceholder")).withBaseSizeTiB(13L)
            .withExtendedCapacitySizeTiB(29L).withPublicNetworkAccess(PublicNetworkAccess.ENABLED)
            .withAutoScaleProperties(new AutoScaleProperties().withScaleUpProperties(new ScaleUpProperties()
                .withUnusedSizeTiB(24L).withIncreaseCapacityUnitByTiB(4L).withCapacityUnitScaleUpLimitTiB(17L)
                .withAutoScalePolicyEnforcement(AutoScalePolicyEnforcement.NONE)))
            .apply();
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
