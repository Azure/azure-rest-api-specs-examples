
import com.azure.resourcemanager.appcontainers.models.Container;
import com.azure.resourcemanager.appcontainers.models.ContainerAppProbe;
import com.azure.resourcemanager.appcontainers.models.ContainerAppProbeHttpGet;
import com.azure.resourcemanager.appcontainers.models.ContainerAppProbeHttpGetHttpHeadersItem;
import com.azure.resourcemanager.appcontainers.models.ContainerResources;
import com.azure.resourcemanager.appcontainers.models.ExtendedLocation;
import com.azure.resourcemanager.appcontainers.models.ExtendedLocationTypes;
import com.azure.resourcemanager.appcontainers.models.InitContainer;
import com.azure.resourcemanager.appcontainers.models.JobConfiguration;
import com.azure.resourcemanager.appcontainers.models.JobConfigurationManualTriggerConfig;
import com.azure.resourcemanager.appcontainers.models.JobTemplate;
import com.azure.resourcemanager.appcontainers.models.TriggerType;
import com.azure.resourcemanager.appcontainers.models.Type;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Jobs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/
     * Job_CreateorUpdate_ConnectedEnvironment.json
     */
    /**
     * Sample code: Create or Update Container Apps Job On A Connected Environment.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void createOrUpdateContainerAppsJobOnAConnectedEnvironment(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.jobs().define("testcontainerAppsJob0").withRegion("East US").withExistingResourceGroup("rg")
            .withExtendedLocation(new ExtendedLocation().withName(
                "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation")
                .withType(ExtendedLocationTypes.CUSTOM_LOCATION))
            .withEnvironmentId(
                "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/connectedEnvironments/demokube")
            .withConfiguration(new JobConfiguration().withTriggerType(TriggerType.MANUAL).withReplicaTimeout(10)
                .withReplicaRetryLimit(10).withManualTriggerConfig(
                    new JobConfigurationManualTriggerConfig().withReplicaCompletionCount(1).withParallelism(4)))
            .withTemplate(new JobTemplate()
                .withInitContainers(Arrays.asList(new InitContainer().withImage("repo/testcontainerAppsJob0:v4")
                    .withName("testinitcontainerAppsJob0").withCommand(Arrays.asList("/bin/sh"))
                    .withArgs(Arrays.asList("-c", "while true; do echo hello; sleep 10;done"))
                    .withResources(new ContainerResources().withCpu(0.2D).withMemory("100Mi"))))
                .withContainers(Arrays
                    .asList(new Container().withImage("repo/testcontainerAppsJob0:v1").withName("testcontainerAppsJob0")
                        .withProbes(Arrays.asList(new ContainerAppProbe()
                            .withHttpGet(new ContainerAppProbeHttpGet()
                                .withHttpHeaders(Arrays.asList(new ContainerAppProbeHttpGetHttpHeadersItem()
                                    .withName("Custom-Header").withValue("Awesome")))
                                .withPath("/health").withPort(8080))
                            .withInitialDelaySeconds(5).withPeriodSeconds(3).withType(Type.LIVENESS))))))
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
