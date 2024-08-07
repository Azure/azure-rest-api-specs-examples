
import com.azure.resourcemanager.appcontainers.models.Container;
import com.azure.resourcemanager.appcontainers.models.ContainerAppProbe;
import com.azure.resourcemanager.appcontainers.models.ContainerAppProbeHttpGet;
import com.azure.resourcemanager.appcontainers.models.ContainerAppProbeHttpGetHttpHeadersItem;
import com.azure.resourcemanager.appcontainers.models.ContainerResources;
import com.azure.resourcemanager.appcontainers.models.InitContainer;
import com.azure.resourcemanager.appcontainers.models.JobConfiguration;
import com.azure.resourcemanager.appcontainers.models.JobConfigurationManualTriggerConfig;
import com.azure.resourcemanager.appcontainers.models.JobTemplate;
import com.azure.resourcemanager.appcontainers.models.TriggerType;
import com.azure.resourcemanager.appcontainers.models.Type;
import java.util.Arrays;

/**
 * Samples for Jobs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/Job_CreateorUpdate.json
     */
    /**
     * Sample code: Create or Update Container Apps Job.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        createOrUpdateContainerAppsJob(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.jobs().define("testcontainerappsjob0").withRegion("East US").withExistingResourceGroup("rg")
            .withEnvironmentId(
                "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube")
            .withConfiguration(new JobConfiguration().withTriggerType(TriggerType.MANUAL).withReplicaTimeout(10)
                .withReplicaRetryLimit(10).withManualTriggerConfig(
                    new JobConfigurationManualTriggerConfig().withReplicaCompletionCount(1).withParallelism(4)))
            .withTemplate(new JobTemplate()
                .withInitContainers(Arrays.asList(new InitContainer().withImage("repo/testcontainerappsjob0:v4")
                    .withName("testinitcontainerAppsJob0").withCommand(Arrays.asList("/bin/sh"))
                    .withArgs(Arrays.asList("-c", "while true; do echo hello; sleep 10;done"))
                    .withResources(new ContainerResources().withCpu(0.5D).withMemory("1Gi"))))
                .withContainers(Arrays
                    .asList(new Container().withImage("repo/testcontainerappsjob0:v1").withName("testcontainerappsjob0")
                        .withProbes(Arrays.asList(new ContainerAppProbe()
                            .withHttpGet(new ContainerAppProbeHttpGet()
                                .withHttpHeaders(Arrays.asList(new ContainerAppProbeHttpGetHttpHeadersItem()
                                    .withName("Custom-Header").withValue("Awesome")))
                                .withPath("/health").withPort(8080))
                            .withInitialDelaySeconds(5).withPeriodSeconds(3).withType(Type.LIVENESS))))))
            .create();
    }
}
