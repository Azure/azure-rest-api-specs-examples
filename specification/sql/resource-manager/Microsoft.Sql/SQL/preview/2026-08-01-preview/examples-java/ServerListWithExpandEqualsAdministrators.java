
/**
 * Samples for Servers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ServerListWithExpandEqualsAdministrators.json
     */
    /**
     * Sample code: List servers with $expand=administrators/activedirectory.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        listServersWithExpandAdministratorsActivedirectory(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServers().list(null, com.azure.core.util.Context.NONE);
    }
}
