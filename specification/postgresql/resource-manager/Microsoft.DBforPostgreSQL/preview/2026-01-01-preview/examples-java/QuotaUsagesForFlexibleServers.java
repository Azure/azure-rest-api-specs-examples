
/**
 * Samples for QuotaUsages List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/QuotaUsagesForFlexibleServers.json
     */
    /**
     * Sample code: List of quota usages for servers.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        listOfQuotaUsagesForServers(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.quotaUsages().list("eastus", com.azure.core.util.Context.NONE);
    }
}
