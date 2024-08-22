
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.machinelearning.models.AmlToken;
import com.azure.resourcemanager.machinelearning.models.CommandJob;
import com.azure.resourcemanager.machinelearning.models.CommandJobLimits;
import com.azure.resourcemanager.machinelearning.models.JobResourceConfiguration;
import com.azure.resourcemanager.machinelearning.models.JobService;
import com.azure.resourcemanager.machinelearning.models.LiteralJobInput;
import com.azure.resourcemanager.machinelearning.models.OutputDeliveryMode;
import com.azure.resourcemanager.machinelearning.models.TensorFlow;
import com.azure.resourcemanager.machinelearning.models.UriFileJobOutput;
import java.io.IOException;
import java.time.Duration;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Jobs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Job/CommandJob/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Command Job.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrUpdateCommandJob(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) throws IOException {
        manager.jobs().define("string").withExistingWorkspace("test-rg", "my-aml-workspace")
            .withProperties(new CommandJob().withDescription("string").withTags(mapOf("string", "string"))
                .withProperties(mapOf("string", "string")).withDisplayName("string").withExperimentName("string")
                .withServices(mapOf("string",
                    new JobService().withJobServiceType("string").withPort(1).withEndpoint("string")
                        .withProperties(mapOf("string", "string"))))
                .withComputeId("string").withIdentity(new AmlToken())
                .withResources(new JobResourceConfiguration().withInstanceCount(1).withInstanceType("string")
                    .withProperties(mapOf("string",
                        SerializerFactory.createDefaultManagementSerializerAdapter().deserialize(
                            "{\"e6b6493e-7d5e-4db3-be1e-306ec641327e\":null}", Object.class, SerializerEncoding.JSON))))
                .withCodeId("fakeTokenPlaceholder").withCommand("string").withEnvironmentId("string")
                .withInputs(mapOf("string", new LiteralJobInput().withDescription("string").withValue("string")))
                .withOutputs(mapOf("string",
                    new UriFileJobOutput().withDescription("string").withUri("string")
                        .withMode(OutputDeliveryMode.READ_WRITE_MOUNT)))
                .withDistribution(new TensorFlow().withWorkerCount(1).withParameterServerCount(1))
                .withLimits(new CommandJobLimits().withTimeout(Duration.parse("PT5M")))
                .withEnvironmentVariables(mapOf("string", "string")))
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
