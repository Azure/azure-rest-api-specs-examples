/** Samples for Users ListByDataBoxEdgeDevice. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/UserGetAllInDevice.json
     */
    /**
     * Sample code: UserGetAllInDevice.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void userGetAllInDevice(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .users()
            .listByDataBoxEdgeDevice(
                "testedgedevice", "GroupForEdgeAutomation", null, com.azure.core.util.Context.NONE);
    }
}
