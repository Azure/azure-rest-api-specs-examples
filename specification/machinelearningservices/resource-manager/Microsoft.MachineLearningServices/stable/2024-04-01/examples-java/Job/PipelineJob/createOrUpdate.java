
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.machinelearning.models.JobService;
import com.azure.resourcemanager.machinelearning.models.LiteralJobInput;
import com.azure.resourcemanager.machinelearning.models.OutputDeliveryMode;
import com.azure.resourcemanager.machinelearning.models.PipelineJob;
import com.azure.resourcemanager.machinelearning.models.UriFileJobOutput;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Jobs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Job/PipelineJob/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Pipeline Job.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrUpdatePipelineJob(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) throws IOException {
        manager.jobs().define("string").withExistingWorkspace("test-rg", "my-aml-workspace")
            .withProperties(new PipelineJob().withDescription("string").withTags(mapOf("string", "string"))
                .withProperties(mapOf("string", "string")).withDisplayName("string").withExperimentName("string")
                .withServices(mapOf("string",
                    new JobService().withJobServiceType("string").withPort(1).withEndpoint("string")
                        .withProperties(mapOf("string", "string"))))
                .withComputeId("string")
                .withSettings(SerializerFactory.createDefaultManagementSerializerAdapter().deserialize("{}",
                    Object.class, SerializerEncoding.JSON))
                .withInputs(mapOf("string", new LiteralJobInput().withDescription("string").withValue("string")))
                .withOutputs(mapOf("string", new UriFileJobOutput().withDescription("string").withUri("string")
                    .withMode(OutputDeliveryMode.UPLOAD))))
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
