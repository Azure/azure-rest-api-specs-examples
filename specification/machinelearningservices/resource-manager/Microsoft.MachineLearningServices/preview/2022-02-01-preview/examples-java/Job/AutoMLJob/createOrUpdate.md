Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-machinelearning_1.0.0-beta.1/sdk/machinelearning/azure-resourcemanager-machinelearning/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.machinelearning.models.AmlToken;
import com.azure.resourcemanager.machinelearning.models.AutoMLJob;
import com.azure.resourcemanager.machinelearning.models.CronSchedule;
import com.azure.resourcemanager.machinelearning.models.ImageClassification;
import com.azure.resourcemanager.machinelearning.models.ImageLimitSettings;
import com.azure.resourcemanager.machinelearning.models.ImageModelDistributionSettingsClassification;
import com.azure.resourcemanager.machinelearning.models.ImageModelSettingsClassification;
import com.azure.resourcemanager.machinelearning.models.ImageVerticalDataSettings;
import com.azure.resourcemanager.machinelearning.models.JobService;
import com.azure.resourcemanager.machinelearning.models.MLTableJobInput;
import com.azure.resourcemanager.machinelearning.models.OutputDeliveryMode;
import com.azure.resourcemanager.machinelearning.models.ResourceConfiguration;
import com.azure.resourcemanager.machinelearning.models.ScheduleStatus;
import com.azure.resourcemanager.machinelearning.models.TrainingDataSettings;
import com.azure.resourcemanager.machinelearning.models.UriFileJobOutput;
import java.io.IOException;
import java.time.OffsetDateTime;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Jobs CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Job/AutoMLJob/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate AutoML Job.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void createOrUpdateAutoMLJob(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) throws IOException {
        manager
            .jobs()
            .define("string")
            .withExistingWorkspace("test-rg", "my-aml-workspace")
            .withProperties(
                new AutoMLJob()
                    .withDescription("string")
                    .withProperties(mapOf("string", "string"))
                    .withTags(mapOf("string", "string"))
                    .withComputeId("string")
                    .withDisplayName("string")
                    .withExperimentName("string")
                    .withIdentity(new AmlToken())
                    .withIsArchived(false)
                    .withSchedule(
                        new CronSchedule()
                            .withEndTime(OffsetDateTime.parse("2020-01-01T12:34:56.999Z"))
                            .withScheduleStatus(ScheduleStatus.DISABLED)
                            .withStartTime(OffsetDateTime.parse("2020-01-01T12:34:56.999Z"))
                            .withTimeZone("string")
                            .withExpression("string"))
                    .withServices(
                        mapOf(
                            "string",
                            new JobService()
                                .withEndpoint("string")
                                .withJobServiceType("string")
                                .withPort(1)
                                .withProperties(mapOf("string", "string"))))
                    .withEnvironmentId("string")
                    .withEnvironmentVariables(mapOf("string", "string"))
                    .withOutputs(
                        mapOf(
                            "string",
                            new UriFileJobOutput()
                                .withDescription("string")
                                .withMode(OutputDeliveryMode.READ_WRITE_MOUNT)
                                .withUri("string")))
                    .withResources(
                        new ResourceConfiguration()
                            .withInstanceCount(1)
                            .withInstanceType("string")
                            .withProperties(
                                mapOf(
                                    "string",
                                    SerializerFactory
                                        .createDefaultManagementSerializerAdapter()
                                        .deserialize(
                                            "{\"9bec0ab0-c62f-4fa9-a97c-7b24bbcc90ad\":null}",
                                            Object.class,
                                            SerializerEncoding.JSON))))
                    .withTaskDetails(
                        new ImageClassification()
                            .withModelSettings(new ImageModelSettingsClassification().withValidationCropSize(2))
                            .withSearchSpace(
                                Arrays
                                    .asList(
                                        new ImageModelDistributionSettingsClassification()
                                            .withValidationCropSize("choice(2, 360)")))
                            .withDataSettings(
                                new ImageVerticalDataSettings()
                                    .withTargetColumnName("string")
                                    .withTrainingData(
                                        new TrainingDataSettings().withData(new MLTableJobInput().withUri("string"))))
                            .withLimitSettings(new ImageLimitSettings().withMaxTrials(2))))
            .create();
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
```
