import com.azure.resourcemanager.cosmos.models.CommandPostBody;

/** Samples for CassandraClusters InvokeCommand. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-04-15/examples/CosmosDBManagedCassandraCommand.json
     */
    /**
     * Sample code: CosmosDBManagedCassandraCommand.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBManagedCassandraCommand(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getCassandraClusters()
            .invokeCommand(
                "cassandra-prod-rg",
                "cassandra-prod",
                new CommandPostBody().withCommand("nodetool status").withHost("10.0.1.12"),
                com.azure.core.util.Context.NONE);
    }
}
