
import com.azure.resourcemanager.postgresqlflexibleserver.models.DbAgentForUpdate;
import com.azure.resourcemanager.postgresqlflexibleserver.models.DbAgentForUpdateProperties;
import com.azure.resourcemanager.postgresqlflexibleserver.models.DbAgentForUpdateState;

/**
 * Samples for DbAgents CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/DBAgentUpdateEnable.json
     */
    /**
     * Sample code: Enable the singleton Default database agent for a server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void enableTheSingletonDefaultDatabaseAgentForAServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.dbAgents().createOrUpdate("exampleresourcegroup", "exampleserver",
            new DbAgentForUpdate()
                .withProperties(new DbAgentForUpdateProperties().withState(DbAgentForUpdateState.ENABLED)),
            com.azure.core.util.Context.NONE);
    }
}
