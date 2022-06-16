import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.machinelearning.models.BatchDeploymentData;
import com.azure.resourcemanager.machinelearning.models.BatchLoggingLevel;
import com.azure.resourcemanager.machinelearning.models.BatchOutputAction;
import com.azure.resourcemanager.machinelearning.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.machinelearning.models.PartialBatchDeployment;
import com.azure.resourcemanager.machinelearning.models.PartialBatchRetrySettings;
import com.azure.resourcemanager.machinelearning.models.PartialCodeConfiguration;
import com.azure.resourcemanager.machinelearning.models.PartialIdAssetReference;
import com.azure.resourcemanager.machinelearning.models.PartialManagedServiceIdentity;
import com.azure.resourcemanager.machinelearning.models.PartialSku;
import com.azure.resourcemanager.machinelearning.models.SkuTier;
import java.io.IOException;
import java.time.Duration;
import java.util.HashMap;
import java.util.Map;

/** Samples for BatchDeployments Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/BatchDeployment/update.json
     */
    /**
     * Sample code: Update Batch Deployment.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void updateBatchDeployment(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) throws IOException {
        BatchDeploymentData resource =
            manager
                .batchDeployments()
                .getWithResponse("test-rg", "my-aml-workspace", "testEndpointName", "testDeploymentName", Context.NONE)
                .getValue();
        resource
            .update()
            .withTags(mapOf())
            .withIdentity(
                new PartialManagedServiceIdentity()
                    .withType(ManagedServiceIdentityType.SYSTEM_ASSIGNED)
                    .withUserAssignedIdentities(
                        mapOf(
                            "string",
                            SerializerFactory
                                .createDefaultManagementSerializerAdapter()
                                .deserialize("{}", Object.class, SerializerEncoding.JSON))))
            .withKind("string")
            .withProperties(
                new PartialBatchDeployment()
                    .withCodeConfiguration(
                        new PartialCodeConfiguration().withCodeId("string").withScoringScript("string"))
                    .withCompute("string")
                    .withDescription("string")
                    .withEnvironmentId("string")
                    .withEnvironmentVariables(mapOf("string", "string"))
                    .withErrorThreshold(1)
                    .withLoggingLevel(BatchLoggingLevel.INFO)
                    .withMaxConcurrencyPerInstance(1)
                    .withMiniBatchSize(1L)
                    .withModel(new PartialIdAssetReference().withAssetId("string"))
                    .withOutputAction(BatchOutputAction.SUMMARY_ONLY)
                    .withOutputFileName("string")
                    .withProperties(mapOf("string", "string"))
                    .withRetrySettings(
                        new PartialBatchRetrySettings().withMaxRetries(1).withTimeout(Duration.parse("PT5M"))))
            .withSku(
                new PartialSku()
                    .withCapacity(1)
                    .withFamily("string")
                    .withName("string")
                    .withSize("string")
                    .withTier(SkuTier.FREE))
            .apply();
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
