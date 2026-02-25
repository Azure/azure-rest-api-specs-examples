
/**
 * Samples for CapabilitiesByServer List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/CapabilitiesByServerList.json
     */
    /**
     * Sample code: List the capabilities available for a given server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void listTheCapabilitiesAvailableForAGivenServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.capabilitiesByServers().list("exampleresourcegroup", "exampleserver", com.azure.core.util.Context.NONE);
    }
}
