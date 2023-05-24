/** Samples for Jobs GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/JobsGetCopyStuck.json
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
                "YourResourceGroupName", "TestJobName1", "details", com.azure.core.util.Context.NONE);
    }
}
