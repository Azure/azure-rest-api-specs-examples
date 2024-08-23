
/**
 * Samples for TimeZones ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedInstanceTimeZoneListByLocation
     * .json
     */
    /**
     * Sample code: List managed instance time zones by location.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listManagedInstanceTimeZonesByLocation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getTimeZones().listByLocation("canadaeast",
            com.azure.core.util.Context.NONE);
    }
}
