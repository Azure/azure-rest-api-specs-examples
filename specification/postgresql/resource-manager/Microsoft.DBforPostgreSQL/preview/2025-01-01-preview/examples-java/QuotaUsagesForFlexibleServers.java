
/**
 * Samples for QuotaUsages List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/
     * QuotaUsagesForFlexibleServers.json
     */
    /**
     * Sample code: List of quota usages for flexible servers.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void listOfQuotaUsagesForFlexibleServers(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.quotaUsages().list("westus", com.azure.core.util.Context.NONE);
    }
}
