
/**
 * Samples for Service Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBGraphAPIComputeServiceDelete.json
     */
    /**
     * Sample code: GraphAPIComputeServiceDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void graphAPIComputeServiceDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getServices().delete("rg1", "ddb1", "GraphAPICompute",
            com.azure.core.util.Context.NONE);
    }
}
