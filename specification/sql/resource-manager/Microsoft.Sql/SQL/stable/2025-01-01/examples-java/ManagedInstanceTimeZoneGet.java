
/**
 * Samples for TimeZones Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceTimeZoneGet.json
     */
    /**
     * Sample code: Get managed instance time zone.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getManagedInstanceTimeZone(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getTimeZones().getWithResponse("canadaeast", "Haiti Standard Time",
            com.azure.core.util.Context.NONE);
    }
}
