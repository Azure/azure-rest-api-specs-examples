
import com.azure.resourcemanager.postgresqlflexibleserver.models.StartMajorVersionUpgradePrecheckRequest;

/**
 * Samples for Servers StartMajorVersionUpgradePrecheck.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/ServersStartMajorVersionUpgradePrecheck.json
     */
    /**
     * Sample code: Start a major version upgrade precheck validation for a PostgreSQL flexible server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void startAMajorVersionUpgradePrecheckValidationForAPostgreSQLFlexibleServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().startMajorVersionUpgradePrecheck("exampleresourcegroup", "exampleserver",
            new StartMajorVersionUpgradePrecheckRequest().withTargetVersion("18"), com.azure.core.util.Context.NONE);
    }
}
