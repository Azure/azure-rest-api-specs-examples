/** Samples for BandwidthSchedules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/BandwidthScheduleGet.json
     */
    /**
     * Sample code: BandwidthScheduleGet.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void bandwidthScheduleGet(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .bandwidthSchedules()
            .getWithResponse(
                "testedgedevice", "bandwidth-1", "GroupForEdgeAutomation", com.azure.core.util.Context.NONE);
    }
}
