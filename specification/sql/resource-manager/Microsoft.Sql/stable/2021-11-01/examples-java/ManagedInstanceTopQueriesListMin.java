
/**
 * Samples for ManagedInstances ListByManagedInstance.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedInstanceTopQueriesListMin.json
     */
    /**
     * Sample code: Obtain list of instance's top resource consuming queries. Minimal request and response.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void obtainListOfInstanceSTopResourceConsumingQueriesMinimalRequestAndResponse(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstances().listByManagedInstance("sqlcrudtest-7398",
            "sqlcrudtest-4645", null, null, null, null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
