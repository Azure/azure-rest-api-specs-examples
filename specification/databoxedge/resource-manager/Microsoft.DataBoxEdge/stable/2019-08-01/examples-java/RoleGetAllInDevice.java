/** Samples for Roles ListByDataBoxEdgeDevice. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/RoleGetAllInDevice.json
     */
    /**
     * Sample code: RoleGetAllInDevice.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void roleGetAllInDevice(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .roles()
            .listByDataBoxEdgeDevice("testedgedevice", "GroupForEdgeAutomation", com.azure.core.util.Context.NONE);
    }
}
