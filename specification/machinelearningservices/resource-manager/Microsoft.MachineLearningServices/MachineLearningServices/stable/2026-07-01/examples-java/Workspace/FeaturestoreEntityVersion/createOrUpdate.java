
import com.azure.resourcemanager.machinelearning.models.FeatureDataType;
import com.azure.resourcemanager.machinelearning.models.FeaturestoreEntityVersionProperties;
import com.azure.resourcemanager.machinelearning.models.IndexColumn;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for FeaturestoreEntityVersions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/FeaturestoreEntityVersion/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Workspace Featurestore Entity Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrUpdateWorkspaceFeaturestoreEntityVersion(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.featurestoreEntityVersions().define("string")
            .withExistingFeaturestoreEntity("test-rg", "my-aml-workspace", "string")
            .withProperties(new FeaturestoreEntityVersionProperties().withDescription("string")
                .withProperties(mapOf("string", "string")).withTags(mapOf("string", "string")).withIsAnonymous(false)
                .withIsArchived(false).withIndexColumns(
                    Arrays.asList(new IndexColumn().withColumnName("string").withDataType(FeatureDataType.DATETIME))))
            .create();
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
