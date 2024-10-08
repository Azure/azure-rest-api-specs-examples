
/**
 * Samples for TimeZones Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedInstanceTimeZoneGet.json
     */
    /**
     * Sample code: Get managed instance time zone.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getManagedInstanceTimeZone(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getTimeZones().getWithResponse("canadaeast", "Haiti Standard Time",
            com.azure.core.util.Context.NONE);
    }
}
