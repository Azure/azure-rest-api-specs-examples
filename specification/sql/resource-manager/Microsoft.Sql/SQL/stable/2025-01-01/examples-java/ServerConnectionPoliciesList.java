
/**
 * Samples for ServerConnectionPolicies ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerConnectionPoliciesList.json
     */
    /**
     * Sample code: Lists a servers connection policies.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listsAServersConnectionPolicies(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerConnectionPolicies().listByServer("rgtest-12", "servertest-6285",
            com.azure.core.util.Context.NONE);
    }
}
