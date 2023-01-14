/** Samples for Devices ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/DataBoxEdgeDeviceGetByResourceGroup.json
     */
    /**
     * Sample code: DataBoxEdgeDeviceGetByResourceGroup.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void dataBoxEdgeDeviceGetByResourceGroup(
        com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager.devices().listByResourceGroup("GroupForEdgeAutomation", null, com.azure.core.util.Context.NONE);
    }
}
