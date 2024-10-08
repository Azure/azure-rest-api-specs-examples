
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.machinelearning.models.BatchDeploymentProperties;
import com.azure.resourcemanager.machinelearning.models.BatchLoggingLevel;
import com.azure.resourcemanager.machinelearning.models.BatchOutputAction;
import com.azure.resourcemanager.machinelearning.models.BatchRetrySettings;
import com.azure.resourcemanager.machinelearning.models.CodeConfiguration;
import com.azure.resourcemanager.machinelearning.models.DeploymentResourceConfiguration;
import com.azure.resourcemanager.machinelearning.models.IdAssetReference;
import com.azure.resourcemanager.machinelearning.models.ManagedServiceIdentity;
import com.azure.resourcemanager.machinelearning.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.machinelearning.models.Sku;
import com.azure.resourcemanager.machinelearning.models.SkuTier;
import com.azure.resourcemanager.machinelearning.models.UserAssignedIdentity;
import java.io.IOException;
import java.time.Duration;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for BatchDeployments CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/BatchDeployment/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Workspace Batch Deployment.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrUpdateWorkspaceBatchDeployment(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) throws IOException {
        manager.batchDeployments().define("testDeploymentName").withRegion("string")
            .withExistingBatchEndpoint("test-rg", "my-aml-workspace", "testEndpointName")
            .withProperties(
                new BatchDeploymentProperties().withDescription("string").withProperties(mapOf("string", "string"))
                    .withCodeConfiguration(
                        new CodeConfiguration().withCodeId("fakeTokenPlaceholder").withScoringScript("string"))
                    .withEnvironmentId("string").withEnvironmentVariables(mapOf("string", "string"))
                    .withCompute("string").withErrorThreshold(1)
                    .withRetrySettings(new BatchRetrySettings().withMaxRetries(1).withTimeout(Duration.parse("PT5M")))
                    .withMiniBatchSize(1L).withLoggingLevel(BatchLoggingLevel.INFO)
                    .withModel(new IdAssetReference().withAssetId("string")).withMaxConcurrencyPerInstance(1)
                    .withOutputAction(BatchOutputAction.SUMMARY_ONLY).withOutputFileName("string")
                    .withResources(new DeploymentResourceConfiguration().withInstanceCount(1).withInstanceType("string")
                        .withProperties(mapOf("string",
                            SerializerFactory.createDefaultManagementSerializerAdapter().deserialize(
                                "{\"cd3c37dc-2876-4ca4-8a54-21bd7619724a\":null}", Object.class,
                                SerializerEncoding.JSON)))))
            .withTags(mapOf()).withKind("string")
            .withIdentity(new ManagedServiceIdentity().withType(ManagedServiceIdentityType.SYSTEM_ASSIGNED)
                .withUserAssignedIdentities(mapOf("string", new UserAssignedIdentity())))
            .withSku(new Sku().withName("string").withTier(SkuTier.FREE).withSize("string").withFamily("string")
                .withCapacity(1))
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
