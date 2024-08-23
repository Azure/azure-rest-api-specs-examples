
/**
 * Samples for ManagedInstances List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedInstanceList.json
     */
    /**
     * Sample code: List managed instances.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listManagedInstances(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstances().list(null, com.azure.core.util.Context.NONE);
    }
}
