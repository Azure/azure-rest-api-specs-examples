/** Samples for Orders Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/OrderDelete.json
     */
    /**
     * Sample code: OrderDelete.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void orderDelete(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager.orders().delete("testedgedevice", "GroupForEdgeAutomation", com.azure.core.util.Context.NONE);
    }
}
