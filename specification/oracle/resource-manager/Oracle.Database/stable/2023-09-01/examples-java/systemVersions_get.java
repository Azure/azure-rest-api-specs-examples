
/**
 * Samples for SystemVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/systemVersions_get.json
     */
    /**
     * Sample code: Get Exadata System Version.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void getExadataSystemVersion(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.systemVersions().getWithResponse("eastus", "22.x", com.azure.core.util.Context.NONE);
    }
}
