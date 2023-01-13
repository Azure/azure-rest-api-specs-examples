/** Samples for Nodes ListByDataBoxEdgeDevice. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/NodeGetAllInDevice.json
     */
    /**
     * Sample code: NodesGetAllInDevice.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void nodesGetAllInDevice(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .nodes()
            .listByDataBoxEdgeDevice("testedgedevice", "GroupForEdgeAutomation", com.azure.core.util.Context.NONE);
    }
}
