
import com.azure.resourcemanager.appcontainers.models.ContainerType;
import com.azure.resourcemanager.appcontainers.models.DynamicPoolConfiguration;
import com.azure.resourcemanager.appcontainers.models.LifecycleConfiguration;
import com.azure.resourcemanager.appcontainers.models.LifecycleType;
import com.azure.resourcemanager.appcontainers.models.PoolManagementType;
import com.azure.resourcemanager.appcontainers.models.ScaleConfiguration;
import com.azure.resourcemanager.appcontainers.models.SessionNetworkConfiguration;
import com.azure.resourcemanager.appcontainers.models.SessionNetworkStatus;

/**
 * Samples for ContainerAppsSessionPools CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/SessionPools_McpServer_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or Update Session Pool with MCP server.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void createOrUpdateSessionPoolWithMCPServer(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsSessionPools().define("testsessionpool").withRegion("East US")
            .withExistingResourceGroup("rg").withContainerType(ContainerType.fromString("Shell"))
            .withPoolManagementType(PoolManagementType.DYNAMIC)
            .withScaleConfiguration(new ScaleConfiguration().withMaxConcurrentSessions(50))
            .withDynamicPoolConfiguration(new DynamicPoolConfiguration().withLifecycleConfiguration(
                new LifecycleConfiguration().withLifecycleType(LifecycleType.TIMED).withCooldownPeriodInSeconds(600)))
            .withSessionNetworkConfiguration(
                new SessionNetworkConfiguration().withStatus(SessionNetworkStatus.EGRESS_ENABLED))
            .create();
    }
}
