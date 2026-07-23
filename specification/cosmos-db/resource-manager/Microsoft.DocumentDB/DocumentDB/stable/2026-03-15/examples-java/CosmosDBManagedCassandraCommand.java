
import com.azure.resourcemanager.cosmos.models.CommandPostBody;

/**
 * Samples for CassandraClusters InvokeCommand.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBManagedCassandraCommand.json
     */
    /**
     * Sample code: CosmosDBManagedCassandraCommand.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBManagedCassandraCommand(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getCassandraClusters().invokeCommand("cassandra-prod-rg", "cassandra-prod",
            new CommandPostBody().withCommand("nodetool status").withHost("10.0.1.12"),
            com.azure.core.util.Context.NONE);
    }
}
