
import com.azure.resourcemanager.elasticsan.models.AutoScalePolicyEnforcement;
import com.azure.resourcemanager.elasticsan.models.AutoScaleProperties;
import com.azure.resourcemanager.elasticsan.models.ElasticSanVersion;
import com.azure.resourcemanager.elasticsan.models.PublicNetworkAccess;
import com.azure.resourcemanager.elasticsan.models.ScaleUpProperties;
import com.azure.resourcemanager.elasticsan.models.Sku;
import com.azure.resourcemanager.elasticsan.models.SkuName;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ElasticSans Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/ElasticSans_V2_Create_MaximumSet_Gen.json
     */
    /**
     * Sample code: ElasticSans_V2_Create_MaximumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void
        elasticSansV2CreateMaximumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.elasticSans().define("elasticsanname").withRegion("France Central")
            .withExistingResourceGroup("resourcegroupname").withSku(new Sku().withName(SkuName.ELASTIC_SAN_LRS))
            .withTags(mapOf("key9316", "fakeTokenPlaceholder")).withAvailabilityZones(Arrays.asList("1"))
            .withTotalIops(22L).withTotalMBps(4L).withTotalSizeTiB(27L)
            .withPublicNetworkAccess(PublicNetworkAccess.ENABLED)
            .withAutoScaleProperties(new AutoScaleProperties().withScaleUpProperties(new ScaleUpProperties()
                .withUnusedSizeTiB(24L).withIncreaseCapacityUnitByTiB(4L).withCapacityUnitScaleUpLimitTiB(17L)
                .withAutoScalePolicyEnforcement(AutoScalePolicyEnforcement.NONE)))
            .withVersion(ElasticSanVersion.V2).create();
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
