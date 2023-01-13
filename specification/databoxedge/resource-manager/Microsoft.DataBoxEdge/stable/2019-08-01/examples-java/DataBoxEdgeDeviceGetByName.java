/** Samples for Devices GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/DataBoxEdgeDeviceGetByName.json
     */
    /**
     * Sample code: DataBoxEdgeDeviceGetByName.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void dataBoxEdgeDeviceGetByName(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .devices()
            .getByResourceGroupWithResponse(
                "GroupForEdgeAutomation", "testedgedevice", com.azure.core.util.Context.NONE);
    }
}
