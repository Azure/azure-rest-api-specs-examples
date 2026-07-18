
/**
 * Samples for TimeZones ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceTimeZoneListByLocation.json
     */
    /**
     * Sample code: List managed instance time zones by location.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listManagedInstanceTimeZonesByLocation(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getTimeZones().listByLocation("canadaeast", com.azure.core.util.Context.NONE);
    }
}
