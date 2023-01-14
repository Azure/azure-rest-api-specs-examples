/** Samples for Devices List. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/DataBoxEdgeDeviceGetBySubscription.json
     */
    /**
     * Sample code: DataBoxEdgeDeviceGetBySubscription.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void dataBoxEdgeDeviceGetBySubscription(
        com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager.devices().list(null, com.azure.core.util.Context.NONE);
    }
}
