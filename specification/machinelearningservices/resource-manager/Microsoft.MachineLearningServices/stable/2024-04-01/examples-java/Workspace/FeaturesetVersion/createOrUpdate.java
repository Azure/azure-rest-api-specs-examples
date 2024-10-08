
import com.azure.resourcemanager.machinelearning.models.EmailNotificationEnableType;
import com.azure.resourcemanager.machinelearning.models.FeaturesetSpecification;
import com.azure.resourcemanager.machinelearning.models.FeaturesetVersionProperties;
import com.azure.resourcemanager.machinelearning.models.MaterializationComputeResource;
import com.azure.resourcemanager.machinelearning.models.MaterializationSettings;
import com.azure.resourcemanager.machinelearning.models.MaterializationStoreType;
import com.azure.resourcemanager.machinelearning.models.NotificationSetting;
import com.azure.resourcemanager.machinelearning.models.RecurrenceFrequency;
import com.azure.resourcemanager.machinelearning.models.RecurrenceSchedule;
import com.azure.resourcemanager.machinelearning.models.RecurrenceTrigger;
import com.azure.resourcemanager.machinelearning.models.WeekDay;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for FeaturesetVersions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/FeaturesetVersion/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Workspace Featureset Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrUpdateWorkspaceFeaturesetVersion(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.featuresetVersions().define("string").withExistingFeatureset("test-rg", "my-aml-workspace", "string")
            .withProperties(new FeaturesetVersionProperties().withDescription("string")
                .withTags(mapOf("string", "string")).withProperties(mapOf("string", "string")).withIsArchived(false)
                .withIsAnonymous(false).withSpecification(new FeaturesetSpecification().withPath("string"))
                .withMaterializationSettings(new MaterializationSettings()
                    .withStoreType(MaterializationStoreType.ONLINE)
                    .withSchedule(new RecurrenceTrigger().withEndTime("string").withStartTime("string")
                        .withTimeZone("string").withFrequency(RecurrenceFrequency.DAY).withInterval(1)
                        .withSchedule(new RecurrenceSchedule().withHours(Arrays.asList(1)).withMinutes(Arrays.asList(1))
                            .withWeekDays(Arrays.asList(WeekDay.MONDAY)).withMonthDays(Arrays.asList(1))))
                    .withNotification(new NotificationSetting().withEmails(Arrays.asList("string"))
                        .withEmailOn(Arrays.asList(EmailNotificationEnableType.JOB_FAILED)))
                    .withResource(new MaterializationComputeResource().withInstanceType("string"))
                    .withSparkConfiguration(mapOf("string", "string")))
                .withStage("string").withEntities(Arrays.asList("string")))
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
