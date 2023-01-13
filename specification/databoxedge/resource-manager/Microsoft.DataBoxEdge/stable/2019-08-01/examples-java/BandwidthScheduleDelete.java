/** Samples for BandwidthSchedules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/BandwidthScheduleDelete.json
     */
    /**
     * Sample code: BandwidthScheduleDelete.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void bandwidthScheduleDelete(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .bandwidthSchedules()
            .delete("testedgedevice", "bandwidth-1", "GroupForEdgeAutomation", com.azure.core.util.Context.NONE);
    }
}
