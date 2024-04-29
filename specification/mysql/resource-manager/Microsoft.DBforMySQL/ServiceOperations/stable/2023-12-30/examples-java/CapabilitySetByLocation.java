
/**
 * Samples for LocationBasedCapabilitySet Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/ServiceOperations/stable/2023-12-30/examples/
     * CapabilitySetByLocation.json
     */
    /**
     * Sample code: CapabilityResult.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void capabilityResult(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.locationBasedCapabilitySets().getWithResponse("WestUS", "default", com.azure.core.util.Context.NONE);
    }
}
