/** Samples for Roles Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/RoleDelete.json
     */
    /**
     * Sample code: RoleDelete.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void roleDelete(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .roles()
            .delete("testedgedevice", "IoTRole1", "GroupForEdgeAutomation", com.azure.core.util.Context.NONE);
    }
}
