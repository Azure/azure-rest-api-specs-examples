/** Samples for Jobs GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2021-03-01/examples/JobsGetCopyStuck.json
     */
    /**
     * Sample code: JobsGetCopyStuck.
     *
     * @param manager Entry point to DataBoxManager.
     */
    public static void jobsGetCopyStuck(com.azure.resourcemanager.databox.DataBoxManager manager) {
        manager
            .jobs()
            .getByResourceGroupWithResponse(
                "dmstestresource", "TJx-637505258985313014", "details", com.azure.core.util.Context.NONE);
    }
}
