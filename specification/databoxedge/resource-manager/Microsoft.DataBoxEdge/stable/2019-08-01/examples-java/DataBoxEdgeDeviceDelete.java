/** Samples for Devices Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/DataBoxEdgeDeviceDelete.json
     */
    /**
     * Sample code: DataBoxEdgeDeviceDelete.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void dataBoxEdgeDeviceDelete(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager.devices().delete("GroupForEdgeAutomation", "testedgedevice", com.azure.core.util.Context.NONE);
    }
}
