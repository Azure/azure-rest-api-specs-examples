
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.machinelearning.models.Goal;
import com.azure.resourcemanager.machinelearning.models.GridSamplingAlgorithm;
import com.azure.resourcemanager.machinelearning.models.JobResourceConfiguration;
import com.azure.resourcemanager.machinelearning.models.JobService;
import com.azure.resourcemanager.machinelearning.models.MedianStoppingPolicy;
import com.azure.resourcemanager.machinelearning.models.Mpi;
import com.azure.resourcemanager.machinelearning.models.Objective;
import com.azure.resourcemanager.machinelearning.models.SweepJob;
import com.azure.resourcemanager.machinelearning.models.SweepJobLimits;
import com.azure.resourcemanager.machinelearning.models.TrialComponent;
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
     * examples/Job/SweepJob/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Sweep Job.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrUpdateSweepJob(com.azure.resourcemanager.machinelearning.MachineLearningManager manager)
        throws IOException {
        manager.jobs().define("string").withExistingWorkspace("test-rg", "my-aml-workspace")
            .withProperties(new SweepJob().withDescription("string").withTags(mapOf("string", "string"))
                .withProperties(mapOf("string", "string")).withDisplayName("string").withExperimentName("string")
                .withServices(mapOf("string",
                    new JobService().withJobServiceType("string").withPort(1).withEndpoint("string")
                        .withProperties(mapOf("string", "string"))))
                .withComputeId("string")
                .withSearchSpace(SerializerFactory.createDefaultManagementSerializerAdapter()
                    .deserialize("{\"string\":{}}", Object.class, SerializerEncoding.JSON))
                .withSamplingAlgorithm(new GridSamplingAlgorithm())
                .withLimits(new SweepJobLimits().withMaxTotalTrials(1).withMaxConcurrentTrials(1)
                    .withTrialTimeout(Duration.parse("PT1S")))
                .withEarlyTermination(new MedianStoppingPolicy().withEvaluationInterval(1).withDelayEvaluation(1))
                .withObjective(new Objective().withPrimaryMetric("string").withGoal(Goal.MINIMIZE))
                .withTrial(new TrialComponent().withCodeId("fakeTokenPlaceholder").withCommand("string")
                    .withEnvironmentId("string").withEnvironmentVariables(mapOf("string", "string"))
                    .withDistribution(new Mpi().withProcessCountPerInstance(1))
                    .withResources(new JobResourceConfiguration().withInstanceCount(1).withInstanceType("string")
                        .withProperties(mapOf("string",
                            SerializerFactory.createDefaultManagementSerializerAdapter().deserialize(
                                "{\"e6b6493e-7d5e-4db3-be1e-306ec641327e\":null}", Object.class,
                                SerializerEncoding.JSON))))))
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
