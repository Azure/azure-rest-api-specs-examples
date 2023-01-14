/** Samples for Alerts Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/AlertGet.json
     */
    /**
     * Sample code: AlertGet.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void alertGet(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .alerts()
            .getWithResponse(
                "testedgedevice",
                "159a00c7-8543-4343-9435-263ac87df3bb",
                "GroupForEdgeAutomation",
                com.azure.core.util.Context.NONE);
    }
}
