/** Samples for Triggers ListByDataBoxEdgeDevice. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/TriggerGetAllInDevice.json
     */
    /**
     * Sample code: TriggerGetAllInDevice.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void triggerGetAllInDevice(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .triggers()
            .listByDataBoxEdgeDevice(
                "testedgedevice", "GroupForEdgeAutomation", null, com.azure.core.util.Context.NONE);
    }
}
