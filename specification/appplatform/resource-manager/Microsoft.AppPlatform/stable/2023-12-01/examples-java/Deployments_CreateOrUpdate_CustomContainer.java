
import com.azure.resourcemanager.appplatform.fluent.models.DeploymentResourceInner;
import com.azure.resourcemanager.appplatform.models.CustomContainer;
import com.azure.resourcemanager.appplatform.models.CustomContainerUserSourceInfo;
import com.azure.resourcemanager.appplatform.models.DeploymentResourceProperties;
import com.azure.resourcemanager.appplatform.models.DeploymentSettings;
import com.azure.resourcemanager.appplatform.models.HttpGetAction;
import com.azure.resourcemanager.appplatform.models.HttpSchemeType;
import com.azure.resourcemanager.appplatform.models.ImageRegistryCredential;
import com.azure.resourcemanager.appplatform.models.Probe;
import com.azure.resourcemanager.appplatform.models.ResourceRequests;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Deployments CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * Deployments_CreateOrUpdate_CustomContainer.json
     */
    /**
     * Sample code: Deployments_CreateOrUpdate_CustomContainer.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deploymentsCreateOrUpdateCustomContainer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getDeployments()
            .createOrUpdate("myResourceGroup", "myservice", "myapp", "mydeployment",
                new DeploymentResourceInner()
                    .withProperties(new DeploymentResourceProperties()
                        .withSource(new CustomContainerUserSourceInfo().withCustomContainer(new CustomContainer()
                            .withServer("myacr.azurecr.io").withContainerImage("myContainerImage:v1")
                            .withCommand(Arrays.asList("/bin/sh"))
                            .withArgs(Arrays.asList("-c", "while true; do echo hello; sleep 10;done"))
                            .withImageRegistryCredential(new ImageRegistryCredential().withUsername("myUsername")
                                .withPassword("fakeTokenPlaceholder"))
                            .withLanguageFramework("springboot")))
                        .withDeploymentSettings(new DeploymentSettings()
                            .withResourceRequests(new ResourceRequests().withCpu("1000m").withMemory("3Gi"))
                            .withEnvironmentVariables(mapOf("env", "test"))
                            .withLivenessProbe(new Probe()
                                .withProbeAction(
                                    new HttpGetAction().withPath("/health").withScheme(HttpSchemeType.HTTP))
                                .withDisableProbe(false).withInitialDelaySeconds(30).withPeriodSeconds(10)
                                .withFailureThreshold(3))
                            .withReadinessProbe(new Probe()
                                .withProbeAction(
                                    new HttpGetAction().withPath("/health").withScheme(HttpSchemeType.HTTP))
                                .withDisableProbe(false).withInitialDelaySeconds(30).withPeriodSeconds(10)
                                .withFailureThreshold(3))
                            .withTerminationGracePeriodSeconds(30))),
                com.azure.core.util.Context.NONE);
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
