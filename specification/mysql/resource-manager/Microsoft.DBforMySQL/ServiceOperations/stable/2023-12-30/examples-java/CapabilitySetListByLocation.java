
/**
 * Samples for LocationBasedCapabilitySet List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/ServiceOperations/stable/2023-12-30/examples/
     * CapabilitySetListByLocation.json
     */
    /**
     * Sample code: CapabilitySetsResult.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void capabilitySetsResult(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.locationBasedCapabilitySets().list("WestUS", com.azure.core.util.Context.NONE);
    }
}
