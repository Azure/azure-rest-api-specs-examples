import com.azure.core.util.Context;

/** Samples for RestorableSqlContainers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBRestorableSqlContainerList.json
     */
    /**
     * Sample code: CosmosDBRestorableSqlContainerList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBRestorableSqlContainerList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getRestorableSqlContainers()
            .list("WestUS", "98a570f2-63db-4117-91f0-366327b7b353", "3fu-hg==", null, null, Context.NONE);
    }
}
