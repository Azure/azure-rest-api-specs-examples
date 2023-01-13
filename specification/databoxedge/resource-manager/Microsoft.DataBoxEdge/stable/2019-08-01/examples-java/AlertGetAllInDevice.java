/** Samples for Alerts ListByDataBoxEdgeDevice. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/AlertGetAllInDevice.json
     */
    /**
     * Sample code: AlertGetAllInDevice.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void alertGetAllInDevice(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .alerts()
            .listByDataBoxEdgeDevice("testedgedevice", "GroupForEdgeAutomation", com.azure.core.util.Context.NONE);
    }
}
