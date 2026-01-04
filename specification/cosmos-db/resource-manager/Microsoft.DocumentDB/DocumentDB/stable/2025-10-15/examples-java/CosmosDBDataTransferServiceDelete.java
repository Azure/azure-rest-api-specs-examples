
/**
 * Samples for Service Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBDataTransferServiceDelete.json
     */
    /**
     * Sample code: DataTransferServiceDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void dataTransferServiceDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getServices().delete("rg1", "ddb1", "DataTransfer",
            com.azure.core.util.Context.NONE);
    }
}
