
/**
 * Samples for LocationBasedCapabilities List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/legacy/stable/2021-05-01/examples/
     * CapabilitiesByLocationList.json
     */
    /**
     * Sample code: CapabilitiesList.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void capabilitiesList(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.locationBasedCapabilities().list("WestUS", com.azure.core.util.Context.NONE);
    }
}
