```java
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.machinelearning.models.Goal;
import com.azure.resourcemanager.machinelearning.models.GridSamplingAlgorithm;
import com.azure.resourcemanager.machinelearning.models.JobService;
import com.azure.resourcemanager.machinelearning.models.MedianStoppingPolicy;
import com.azure.resourcemanager.machinelearning.models.Mpi;
import com.azure.resourcemanager.machinelearning.models.Objective;
import com.azure.resourcemanager.machinelearning.models.ResourceConfiguration;
import com.azure.resourcemanager.machinelearning.models.SweepJob;
import com.azure.resourcemanager.machinelearning.models.SweepJobLimits;
import com.azure.resourcemanager.machinelearning.models.TrialComponent;
import java.io.IOException;
import java.time.Duration;
import java.util.HashMap;
import java.util.Map;

/** Samples for Jobs CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Job/SweepJob/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Sweep Job.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void createOrUpdateSweepJob(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) throws IOException {
        manager
            .jobs()
            .define("string")
            .withExistingWorkspace("test-rg", "my-aml-workspace")
            .withProperties(
                new SweepJob()
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
                    .withEarlyTermination(new MedianStoppingPolicy().withDelayEvaluation(1).withEvaluationInterval(1))
                    .withLimits(
                        new SweepJobLimits()
                            .withMaxConcurrentTrials(1)
                            .withMaxTotalTrials(1)
                            .withTrialTimeout(Duration.parse("PT1S")))
                    .withObjective(new Objective().withGoal(Goal.MINIMIZE).withPrimaryMetric("string"))
                    .withSamplingAlgorithm(new GridSamplingAlgorithm())
                    .withSearchSpace(
                        SerializerFactory
                            .createDefaultManagementSerializerAdapter()
                            .deserialize("{\"string\":{}}", Object.class, SerializerEncoding.JSON))
                    .withTrial(
                        new TrialComponent()
                            .withCodeId("string")
                            .withCommand("string")
                            .withDistribution(new Mpi().withProcessCountPerInstance(1))
                            .withEnvironmentId("string")
                            .withEnvironmentVariables(mapOf("string", "string"))
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
                                                    SerializerEncoding.JSON))))))
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-machinelearning_1.0.0-beta.1/sdk/machinelearning/azure-resourcemanager-machinelearning/README.md) on how to add the SDK to your project and authenticate.
