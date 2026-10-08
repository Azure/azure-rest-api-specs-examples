
import com.azure.resourcemanager.elasticsan.models.QualityOfService;
import com.azure.resourcemanager.elasticsan.models.StorageTargetType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for VolumeGroups Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/VolumeGroups_GeneralPurpose_Create_MinimumSet_Gen.json
     */
    /**
     * Sample code: VolumeGroups_GeneralPurpose_Create_MinimumSet_Gen.
     * 
     * @param manager Entry point to ElasticSanManager.
     */
    public static void
        volumeGroupsGeneralPurposeCreateMinimumSetGen(com.azure.resourcemanager.elasticsan.ElasticSanManager manager) {
        manager.volumeGroups().define("volumegroupname").withExistingElasticSan("resourcegroupname", "elasticsanname")
            .withProtocolType(StorageTargetType.ISCSI).withQualityOfService(QualityOfService.GENERAL_PURPOSE).create();
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
