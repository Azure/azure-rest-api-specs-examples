
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.machinelearning.models.AmlToken;
import com.azure.resourcemanager.machinelearning.models.AutoMLJob;
import com.azure.resourcemanager.machinelearning.models.ImageClassification;
import com.azure.resourcemanager.machinelearning.models.ImageLimitSettings;
import com.azure.resourcemanager.machinelearning.models.ImageModelDistributionSettingsClassification;
import com.azure.resourcemanager.machinelearning.models.ImageModelSettingsClassification;
import com.azure.resourcemanager.machinelearning.models.JobResourceConfiguration;
import com.azure.resourcemanager.machinelearning.models.JobService;
import com.azure.resourcemanager.machinelearning.models.MLTableJobInput;
import com.azure.resourcemanager.machinelearning.models.OutputDeliveryMode;
import com.azure.resourcemanager.machinelearning.models.UriFileJobOutput;
import java.io.IOException;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Jobs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Job/AutoMLJob/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate AutoML Job.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrUpdateAutoMLJob(com.azure.resourcemanager.machinelearning.MachineLearningManager manager)
        throws IOException {
        manager.jobs().define("string").withExistingWorkspace("test-rg", "my-aml-workspace")
            .withProperties(new AutoMLJob().withDescription("string").withTags(mapOf("string", "string"))
                .withProperties(mapOf("string", "string")).withDisplayName("string").withExperimentName("string")
                .withServices(mapOf("string",
                    new JobService().withJobServiceType("string").withPort(1).withEndpoint("string")
                        .withProperties(mapOf("string", "string"))))
                .withComputeId("string").withIsArchived(false).withIdentity(new AmlToken())
                .withResources(new JobResourceConfiguration().withInstanceCount(1).withInstanceType("string")
                    .withProperties(mapOf("string",
                        SerializerFactory.createDefaultManagementSerializerAdapter().deserialize(
                            "{\"9bec0ab0-c62f-4fa9-a97c-7b24bbcc90ad\":null}", Object.class, SerializerEncoding.JSON))))
                .withEnvironmentId("string").withEnvironmentVariables(mapOf("string", "string"))
                .withTaskDetails(new ImageClassification().withTrainingData(new MLTableJobInput().withUri("string"))
                    .withTargetColumnName("string")
                    .withModelSettings(new ImageModelSettingsClassification().withValidationCropSize(2))
                    .withSearchSpace(Arrays.asList(
                        new ImageModelDistributionSettingsClassification().withValidationCropSize("choice(2, 360)")))
                    .withLimitSettings(new ImageLimitSettings().withMaxTrials(2)))
                .withOutputs(mapOf("string", new UriFileJobOutput().withDescription("string").withUri("string")
                    .withMode(OutputDeliveryMode.READ_WRITE_MOUNT))))
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
