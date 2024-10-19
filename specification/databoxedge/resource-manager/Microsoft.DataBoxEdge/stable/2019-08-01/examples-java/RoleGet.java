
/**
 * Samples for Roles Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/RoleGet.json
     */
    /**
     * Sample code: RoleGet.
     * 
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void roleGet(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager.roles().getWithResponse("testedgedevice", "IoTRole1", "GroupForEdgeAutomation",
            com.azure.core.util.Context.NONE);
    }
}
