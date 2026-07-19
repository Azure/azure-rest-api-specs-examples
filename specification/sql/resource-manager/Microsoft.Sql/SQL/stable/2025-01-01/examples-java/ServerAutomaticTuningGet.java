
/**
 * Samples for ServerAutomaticTuning Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerAutomaticTuningGet.json
     */
    /**
     * Sample code: Get a server's automatic tuning settings.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getAServerSAutomaticTuningSettings(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerAutomaticTunings().getWithResponse("default-sql-onebox", "testsvr11",
            com.azure.core.util.Context.NONE);
    }
}
