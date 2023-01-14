/** Samples for Triggers Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/TriggerGet.json
     */
    /**
     * Sample code: TriggerGet.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void triggerGet(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .triggers()
            .getWithResponse("testedgedevice", "trigger1", "GroupForEdgeAutomation", com.azure.core.util.Context.NONE);
    }
}
