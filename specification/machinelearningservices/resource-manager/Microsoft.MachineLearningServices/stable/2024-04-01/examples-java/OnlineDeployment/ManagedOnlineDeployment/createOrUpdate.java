
import com.azure.resourcemanager.machinelearning.models.CodeConfiguration;
import com.azure.resourcemanager.machinelearning.models.DefaultScaleSettings;
import com.azure.resourcemanager.machinelearning.models.ManagedOnlineDeployment;
import com.azure.resourcemanager.machinelearning.models.ManagedServiceIdentity;
import com.azure.resourcemanager.machinelearning.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.machinelearning.models.OnlineRequestSettings;
import com.azure.resourcemanager.machinelearning.models.ProbeSettings;
import com.azure.resourcemanager.machinelearning.models.Sku;
import com.azure.resourcemanager.machinelearning.models.SkuTier;
import com.azure.resourcemanager.machinelearning.models.UserAssignedIdentity;
import java.time.Duration;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for OnlineDeployments CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/OnlineDeployment/ManagedOnlineDeployment/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Managed Online Deployment.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrUpdateManagedOnlineDeployment(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.onlineDeployments().define("testDeploymentName").withRegion("string")
            .withExistingOnlineEndpoint("test-rg", "my-aml-workspace", "testEndpointName")
            .withProperties(
                new ManagedOnlineDeployment().withDescription("string").withProperties(mapOf("string", "string"))
                    .withCodeConfiguration(
                        new CodeConfiguration().withCodeId("fakeTokenPlaceholder").withScoringScript("string"))
                    .withEnvironmentId("string").withEnvironmentVariables(mapOf("string", "string"))
                    .withScaleSettings(new DefaultScaleSettings())
                    .withRequestSettings(new OnlineRequestSettings().withMaxQueueWait(Duration.parse("PT5M"))
                        .withRequestTimeout(Duration.parse("PT5M")).withMaxConcurrentRequestsPerInstance(1))
                    .withModelMountPath("string").withAppInsightsEnabled(false)
                    .withLivenessProbe(new ProbeSettings().withFailureThreshold(1).withSuccessThreshold(1)
                        .withTimeout(Duration.parse("PT5M")).withPeriod(Duration.parse("PT5M"))
                        .withInitialDelay(Duration.parse("PT5M")))
                    .withReadinessProbe(new ProbeSettings().withFailureThreshold(30).withSuccessThreshold(1)
                        .withTimeout(Duration.parse("PT2S")).withPeriod(Duration.parse("PT10S"))
                        .withInitialDelay(Duration.parse("PT1S")))
                    .withInstanceType("string").withModel("string"))
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
