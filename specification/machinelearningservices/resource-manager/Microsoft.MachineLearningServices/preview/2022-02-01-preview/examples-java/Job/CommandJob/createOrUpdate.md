Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-machinelearning_1.0.0-beta.2/sdk/machinelearning/azure-resourcemanager-machinelearning/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.machinelearning.models.AmlToken;
import com.azure.resourcemanager.machinelearning.models.CommandJob;
import com.azure.resourcemanager.machinelearning.models.CommandJobLimits;
import com.azure.resourcemanager.machinelearning.models.JobService;
import com.azure.resourcemanager.machinelearning.models.LiteralJobInput;
import com.azure.resourcemanager.machinelearning.models.OutputDeliveryMode;
import com.azure.resourcemanager.machinelearning.models.ResourceConfiguration;
import com.azure.resourcemanager.machinelearning.models.TensorFlow;
import com.azure.resourcemanager.machinelearning.models.UriFileJobOutput;
import java.io.IOException;
import java.time.Duration;
import java.util.HashMap;
import java.util.Map;

/** Samples for Jobs CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Job/CommandJob/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Command Job.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrUpdateCommandJob(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) throws IOException {
        manager
            .jobs()
            .define("string")
            .withExistingWorkspace("test-rg", "my-aml-workspace")
            .withProperties(
                new CommandJob()
                    .withDescription("string")
                    .withProperties(mapOf("string", "string"))
                    .withTags(mapOf("string", "string"))
                    .withComputeId("string")
                    .withDisplayName("string")
                    .withExperimentName("string")
                    .withIdentity(new AmlToken())
                    .withServices(
                        mapOf(
                            "string",
                            new JobService()
                                .withEndpoint("string")
                                .withJobServiceType("string")
                                .withPort(1)
                                .withProperties(mapOf("string", "string"))))
                    .withCodeId("string")
                    .withCommand("string")
                    .withDistribution(new TensorFlow().withParameterServerCount(1).withWorkerCount(1))
                    .withEnvironmentId("string")
                    .withEnvironmentVariables(mapOf("string", "string"))
                    .withInputs(mapOf("string", new LiteralJobInput().withDescription("string").withValue("string")))
                    .withLimits(new CommandJobLimits().withTimeout(Duration.parse("PT5M")))
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
                                            "{\"e6b6493e-7d5e-4db3-be1e-306ec641327e\":null}",
                                            Object.class,
                                            SerializerEncoding.JSON)))))
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
