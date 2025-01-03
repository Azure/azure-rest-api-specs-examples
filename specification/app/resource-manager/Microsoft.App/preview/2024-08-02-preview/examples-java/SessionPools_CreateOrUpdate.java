
import com.azure.resourcemanager.appcontainers.models.ContainerType;
import com.azure.resourcemanager.appcontainers.models.CustomContainerTemplate;
import com.azure.resourcemanager.appcontainers.models.DynamicPoolConfiguration;
import com.azure.resourcemanager.appcontainers.models.ExecutionType;
import com.azure.resourcemanager.appcontainers.models.PoolManagementType;
import com.azure.resourcemanager.appcontainers.models.ScaleConfiguration;
import com.azure.resourcemanager.appcontainers.models.SessionContainer;
import com.azure.resourcemanager.appcontainers.models.SessionContainerResources;
import com.azure.resourcemanager.appcontainers.models.SessionIngress;
import com.azure.resourcemanager.appcontainers.models.SessionNetworkConfiguration;
import com.azure.resourcemanager.appcontainers.models.SessionNetworkStatus;
import java.util.Arrays;

/**
 * Samples for ContainerAppsSessionPools CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/SessionPools_CreateOrUpdate.
     * json
     */
    /**
     * Sample code: Create or Update Session Pool.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        createOrUpdateSessionPool(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsSessionPools().define("testsessionpool").withRegion("East US")
            .withExistingResourceGroup("rg")
            .withEnvironmentId(
                "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube")
            .withContainerType(ContainerType.CUSTOM_CONTAINER).withPoolManagementType(PoolManagementType.DYNAMIC)
            .withScaleConfiguration(
                new ScaleConfiguration().withMaxConcurrentSessions(500).withReadySessionInstances(100))
            .withDynamicPoolConfiguration(
                new DynamicPoolConfiguration().withExecutionType(ExecutionType.TIMED).withCooldownPeriodInSeconds(600))
            .withCustomContainerTemplate(new CustomContainerTemplate()
                .withContainers(Arrays.asList(new SessionContainer().withImage("repo/testcontainer:v4")
                    .withName("testinitcontainer").withCommand(Arrays.asList("/bin/sh"))
                    .withArgs(Arrays.asList("-c", "while true; do echo hello; sleep 10;done"))
                    .withResources(new SessionContainerResources().withCpu(0.25D).withMemory("0.5Gi"))))
                .withIngress(new SessionIngress().withTargetPort(80)))
            .withSessionNetworkConfiguration(
                new SessionNetworkConfiguration().withStatus(SessionNetworkStatus.EGRESS_ENABLED))
            .create();
    }
}
