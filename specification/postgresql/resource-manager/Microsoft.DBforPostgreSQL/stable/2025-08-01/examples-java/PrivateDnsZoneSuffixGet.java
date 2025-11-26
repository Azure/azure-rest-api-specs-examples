
/**
 * Samples for PrivateDnsZoneSuffix Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/
     * PrivateDnsZoneSuffixGet.json
     */
    /**
     * Sample code: Get the private DNS suffix.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        getThePrivateDNSSuffix(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.privateDnsZoneSuffixes().getWithResponse(com.azure.core.util.Context.NONE);
    }
}
