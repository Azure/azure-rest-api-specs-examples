
import com.azure.resourcemanager.appcontainers.models.Container;
import com.azure.resourcemanager.appcontainers.models.ContainerAppProbe;
import com.azure.resourcemanager.appcontainers.models.ContainerAppProbeHttpGet;
import com.azure.resourcemanager.appcontainers.models.ContainerAppProbeHttpGetHttpHeadersItem;
import com.azure.resourcemanager.appcontainers.models.ContainerResources;
import com.azure.resourcemanager.appcontainers.models.IdentitySettings;
import com.azure.resourcemanager.appcontainers.models.IdentitySettingsLifeCycle;
import com.azure.resourcemanager.appcontainers.models.InitContainer;
import com.azure.resourcemanager.appcontainers.models.JobConfiguration;
import com.azure.resourcemanager.appcontainers.models.JobConfigurationManualTriggerConfig;
import com.azure.resourcemanager.appcontainers.models.JobTemplate;
import com.azure.resourcemanager.appcontainers.models.ManagedServiceIdentity;
import com.azure.resourcemanager.appcontainers.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.appcontainers.models.TriggerType;
import com.azure.resourcemanager.appcontainers.models.Type;
import com.azure.resourcemanager.appcontainers.models.UserAssignedIdentity;
import com.azure.resourcemanager.appcontainers.models.VolumeMount;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Jobs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/Job_CreateorUpdate.json
     */
    /**
     * Sample code: Create or Update Container Apps Job.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        createOrUpdateContainerAppsJob(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.jobs().define("testcontainerAppsJob0").withRegion("East US").withExistingResourceGroup("rg")
            .withIdentity(new ManagedServiceIdentity()
                .withType(ManagedServiceIdentityType.SYSTEM_ASSIGNED_USER_ASSIGNED)
                .withUserAssignedIdentities(mapOf(
                    "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity",
                    new UserAssignedIdentity())))
            .withEnvironmentId(
                "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube")
            .withConfiguration(new JobConfiguration().withTriggerType(TriggerType.MANUAL).withReplicaTimeout(10)
                .withReplicaRetryLimit(10)
                .withManualTriggerConfig(
                    new JobConfigurationManualTriggerConfig().withReplicaCompletionCount(1).withParallelism(4))
                .withIdentitySettings(Arrays.asList(new IdentitySettings().withIdentity(
                    "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity")
                    .withLifecycle(IdentitySettingsLifeCycle.ALL),
                    new IdentitySettings().withIdentity("system").withLifecycle(IdentitySettingsLifeCycle.INIT))))
            .withTemplate(
                new JobTemplate()
                    .withInitContainers(Arrays.asList(new InitContainer().withImage("repo/testcontainerAppsJob0:v4")
                        .withName("testinitcontainerAppsJob0").withCommand(Arrays.asList("/bin/sh"))
                        .withArgs(Arrays.asList("-c", "while true; do echo hello; sleep 10;done"))
                        .withResources(new ContainerResources().withCpu(0.2D).withMemory("100Mi"))))
                    .withContainers(Arrays.asList(
                        new Container().withImage("repo/testcontainerAppsJob0:v1").withName("testcontainerAppsJob0")
                            .withVolumeMounts(Arrays.asList(
                                new VolumeMount().withVolumeName("azurefile").withMountPath("/mnt/path1")
                                    .withSubPath("subPath1"),
                                new VolumeMount().withVolumeName("nfsazurefile").withMountPath("/mnt/path2")
                                    .withSubPath("subPath2")))
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
