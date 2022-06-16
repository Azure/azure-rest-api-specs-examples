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

/** Samples for OnlineDeployments CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/OnlineDeployment/ManagedOnlineDeployment/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Managed Online Deployment.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void createOrUpdateManagedOnlineDeployment(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .onlineDeployments()
            .define("testDeploymentName")
            .withRegion("string")
            .withExistingOnlineEndpoint("test-rg", "my-aml-workspace", "testEndpointName")
            .withProperties(
                new ManagedOnlineDeployment()
                    .withCodeConfiguration(new CodeConfiguration().withCodeId("string").withScoringScript("string"))
                    .withDescription("string")
                    .withEnvironmentId("string")
                    .withEnvironmentVariables(mapOf("string", "string"))
                    .withProperties(mapOf("string", "string"))
                    .withAppInsightsEnabled(false)
                    .withInstanceType("string")
                    .withLivenessProbe(
                        new ProbeSettings()
                            .withFailureThreshold(1)
                            .withInitialDelay(Duration.parse("PT5M"))
                            .withPeriod(Duration.parse("PT5M"))
                            .withSuccessThreshold(1)
                            .withTimeout(Duration.parse("PT5M")))
                    .withModel("string")
                    .withModelMountPath("string")
                    .withReadinessProbe(
                        new ProbeSettings()
                            .withFailureThreshold(30)
                            .withInitialDelay(Duration.parse("PT1S"))
                            .withPeriod(Duration.parse("PT10S"))
                            .withSuccessThreshold(1)
                            .withTimeout(Duration.parse("PT2S")))
                    .withRequestSettings(
                        new OnlineRequestSettings()
                            .withMaxConcurrentRequestsPerInstance(1)
                            .withMaxQueueWait(Duration.parse("PT5M"))
                            .withRequestTimeout(Duration.parse("PT5M")))
                    .withScaleSettings(new DefaultScaleSettings()))
            .withTags(mapOf())
            .withIdentity(
                new ManagedServiceIdentity()
                    .withType(ManagedServiceIdentityType.SYSTEM_ASSIGNED)
                    .withUserAssignedIdentities(mapOf("string", new UserAssignedIdentity())))
            .withKind("string")
            .withSku(
                new Sku()
                    .withName("string")
                    .withTier(SkuTier.FREE)
                    .withSize("string")
                    .withFamily("string")
                    .withCapacity(1))
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
