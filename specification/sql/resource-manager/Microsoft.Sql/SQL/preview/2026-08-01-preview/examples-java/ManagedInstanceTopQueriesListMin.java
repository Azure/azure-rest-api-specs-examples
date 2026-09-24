
/**
 * Samples for ManagedInstances ListByManagedInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ManagedInstanceTopQueriesListMin.json
     */
    /**
     * Sample code: Obtain list of instance's top resource consuming queries. Minimal request and response.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void obtainListOfInstanceSTopResourceConsumingQueriesMinimalRequestAndResponse(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstances().listByManagedInstance("sqlcrudtest-7398", "sqlcrudtest-4645",
            null, null, null, null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
