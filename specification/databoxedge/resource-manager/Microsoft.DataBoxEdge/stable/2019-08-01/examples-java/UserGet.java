/** Samples for Users Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/UserGet.json
     */
    /**
     * Sample code: UserGet.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void userGet(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .users()
            .getWithResponse("testedgedevice", "user1", "GroupForEdgeAutomation", com.azure.core.util.Context.NONE);
    }
}
