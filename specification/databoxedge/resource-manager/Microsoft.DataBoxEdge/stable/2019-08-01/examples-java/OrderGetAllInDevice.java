/** Samples for Orders ListByDataBoxEdgeDevice. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/OrderGetAllInDevice.json
     */
    /**
     * Sample code: OrderGetAllInDevice.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void orderGetAllInDevice(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .orders()
            .listByDataBoxEdgeDevice("testedgedevice", "GroupForEdgeAutomation", com.azure.core.util.Context.NONE);
    }
}
