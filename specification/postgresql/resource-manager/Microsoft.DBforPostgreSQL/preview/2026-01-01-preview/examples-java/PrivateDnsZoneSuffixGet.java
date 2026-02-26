
/**
 * Samples for PrivateDnsZoneSuffix Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/PrivateDnsZoneSuffixGet.json
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
