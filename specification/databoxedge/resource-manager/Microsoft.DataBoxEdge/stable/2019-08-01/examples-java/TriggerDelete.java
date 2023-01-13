/** Samples for Triggers Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/TriggerDelete.json
     */
    /**
     * Sample code: TriggerDelete.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void triggerDelete(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .triggers()
            .delete("testedgedevice", "trigger1", "GroupForEdgeAutomation", com.azure.core.util.Context.NONE);
    }
}
