
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ListOperations.json
     */
    /**
     * Sample code: Lists all of the available SQL Rest API operations.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listsAllOfTheAvailableSQLRestAPIOperations(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getOperations().list(com.azure.core.util.Context.NONE);
    }
}
