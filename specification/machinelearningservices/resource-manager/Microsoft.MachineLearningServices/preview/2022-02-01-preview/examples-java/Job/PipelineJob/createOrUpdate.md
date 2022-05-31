Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-machinelearning_1.0.0-beta.1/sdk/machinelearning/azure-resourcemanager-machinelearning/README.md) on how to add the SDK to your project and authenticate.

```java
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

/** Samples for Jobs CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Job/PipelineJob/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Pipeline Job.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void createOrUpdatePipelineJob(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) throws IOException {
        manager
            .jobs()
            .define("string")
            .withExistingWorkspace("test-rg", "my-aml-workspace")
            .withProperties(
                new PipelineJob()
                    .withDescription("string")
                    .withProperties(mapOf("string", "string"))
                    .withTags(mapOf("string", "string"))
                    .withComputeId("string")
                    .withDisplayName("string")
                    .withExperimentName("string")
                    .withServices(
                        mapOf(
                            "string",
                            new JobService()
                                .withEndpoint("string")
                                .withJobServiceType("string")
                                .withPort(1)
                                .withProperties(mapOf("string", "string"))))
                    .withInputs(mapOf("string", new LiteralJobInput().withDescription("string").withValue("string")))
                    .withOutputs(
                        mapOf(
                            "string",
                            new UriFileJobOutput()
                                .withDescription("string")
                                .withMode(OutputDeliveryMode.UPLOAD)
                                .withUri("string")))
                    .withSettings(
                        SerializerFactory
                            .createDefaultManagementSerializerAdapter()
                            .deserialize("{}", Object.class, SerializerEncoding.JSON)))
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
