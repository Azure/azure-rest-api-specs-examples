/** Samples for Devices GetExtendedInformation. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/ExtendedInfoPost.json
     */
    /**
     * Sample code: ExtendedInfoPost.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void extendedInfoPost(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .devices()
            .getExtendedInformationWithResponse(
                "testedgedevice", "GroupForEdgeAutomation", com.azure.core.util.Context.NONE);
    }
}
