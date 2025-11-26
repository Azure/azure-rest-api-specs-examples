
import com.azure.resourcemanager.postgresqlflexibleserver.models.FailoverMode;
import com.azure.resourcemanager.postgresqlflexibleserver.models.RestartParameter;

/**
 * Samples for Servers Restart.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/
     * ServersRestartWithFailover.json
     */
    /**
     * Sample code: Restart PostgreSQL database engine in a server with a forced failover to standby server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void restartPostgreSQLDatabaseEngineInAServerWithAForcedFailoverToStandbyServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().restart("exampleresourcegroup", "exampleserver",
            new RestartParameter().withRestartWithFailover(true).withFailoverMode(FailoverMode.FORCED_FAILOVER),
            com.azure.core.util.Context.NONE);
    }
}
