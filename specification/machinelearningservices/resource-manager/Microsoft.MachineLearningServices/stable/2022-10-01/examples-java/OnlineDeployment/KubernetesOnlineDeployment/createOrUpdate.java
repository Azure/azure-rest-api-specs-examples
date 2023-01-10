import com.azure.resourcemanager.machinelearning.models.CodeConfiguration;
import com.azure.resourcemanager.machinelearning.models.ContainerResourceRequirements;
import com.azure.resourcemanager.machinelearning.models.ContainerResourceSettings;
import com.azure.resourcemanager.machinelearning.models.DefaultScaleSettings;
import com.azure.resourcemanager.machinelearning.models.KubernetesOnlineDeployment;
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
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/OnlineDeployment/KubernetesOnlineDeployment/createOrUpdate.json
     */
    /**
     * Sample code: CreateOrUpdate Kubernetes Online Deployment.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrUpdateKubernetesOnlineDeployment(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .onlineDeployments()
            .define("testDeploymentName")
            .withRegion("string")
            .withExistingOnlineEndpoint("test-rg", "my-aml-workspace", "testEndpointName")
            .withProperties(
                new KubernetesOnlineDeployment()
                    .withCodeConfiguration(
                        new CodeConfiguration().withCodeId("fakeTokenPlaceholder").withScoringScript("string"))
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
                    .withRequestSettings(
                        new OnlineRequestSettings()
                            .withMaxConcurrentRequestsPerInstance(1)
                            .withMaxQueueWait(Duration.parse("PT5M"))
                            .withRequestTimeout(Duration.parse("PT5M")))
                    .withScaleSettings(new DefaultScaleSettings())
                    .withContainerResourceRequirements(
                        new ContainerResourceRequirements()
                            .withContainerResourceLimits(
                                new ContainerResourceSettings().withCpu("\"1\"").withGpu("\"1\"").withMemory("\"2Gi\""))
                            .withContainerResourceRequests(
                                new ContainerResourceSettings()
                                    .withCpu("\"1\"")
                                    .withGpu("\"1\"")
                                    .withMemory("\"2Gi\""))))
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
