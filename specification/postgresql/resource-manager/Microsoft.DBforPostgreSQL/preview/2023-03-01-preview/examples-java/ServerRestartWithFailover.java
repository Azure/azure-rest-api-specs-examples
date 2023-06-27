import com.azure.resourcemanager.postgresqlflexibleserver.models.FailoverMode;
import com.azure.resourcemanager.postgresqlflexibleserver.models.RestartParameter;

/** Samples for Servers Restart. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-01-preview/examples/ServerRestartWithFailover.json
     */
    /**
     * Sample code: ServerRestartWithFailover.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void serverRestartWithFailover(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager
            .servers()
            .restart(
                "testrg",
                "testserver",
                new RestartParameter().withRestartWithFailover(true).withFailoverMode(FailoverMode.FORCED_FAILOVER),
                com.azure.core.util.Context.NONE);
    }
}
