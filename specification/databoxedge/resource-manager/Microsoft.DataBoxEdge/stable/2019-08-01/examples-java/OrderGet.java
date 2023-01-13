/** Samples for Orders Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/OrderGet.json
     */
    /**
     * Sample code: OrderGet.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void orderGet(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager.orders().getWithResponse("testedgedevice", "GroupForEdgeAutomation", com.azure.core.util.Context.NONE);
    }
}
