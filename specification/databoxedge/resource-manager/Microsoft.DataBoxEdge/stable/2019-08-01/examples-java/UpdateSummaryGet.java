/** Samples for Devices GetUpdateSummary. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/UpdateSummaryGet.json
     */
    /**
     * Sample code: UpdateSummaryGet.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void updateSummaryGet(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .devices()
            .getUpdateSummaryWithResponse("testedgedevice", "GroupForEdgeAutomation", com.azure.core.util.Context.NONE);
    }
}
