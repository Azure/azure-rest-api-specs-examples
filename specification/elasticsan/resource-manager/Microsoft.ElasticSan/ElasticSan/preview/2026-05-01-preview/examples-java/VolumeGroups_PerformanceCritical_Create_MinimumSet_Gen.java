
import com.azure.resourcemanager.elasticsan.models.QualityOfService;
import com.azure.resourcemanager.elasticsan.models.StorageTargetType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for VolumeGroups Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/VolumeGroups_PerformanceCritical_Create_MinimumSet_Gen.json
     */
    /**
     * Sample code: VolumeGroups_PerformanceCritical_Create_MinimumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void volumeGroupsPerformanceCriticalCreateMinimumSetGen(
        com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.volumeGroups().define("volumegroupname").withExistingElasticSan("resourcegroupname", "elasticsanname")
            .withProtocolType(StorageTargetType.DIRECT_ATTACH).withReservedIops(10000).withReservedMBps(800)
            .withQualityOfService(QualityOfService.PERFORMANCE_CRITICAL).create();
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
