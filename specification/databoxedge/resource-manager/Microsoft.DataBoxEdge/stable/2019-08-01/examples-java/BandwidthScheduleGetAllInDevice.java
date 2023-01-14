/** Samples for BandwidthSchedules ListByDataBoxEdgeDevice. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/BandwidthScheduleGetAllInDevice.json
     */
    /**
     * Sample code: BandwidthScheduleGetAllInDevice.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void bandwidthScheduleGetAllInDevice(
        com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .bandwidthSchedules()
            .listByDataBoxEdgeDevice("testedgedevice", "GroupForEdgeAutomation", com.azure.core.util.Context.NONE);
    }
}
