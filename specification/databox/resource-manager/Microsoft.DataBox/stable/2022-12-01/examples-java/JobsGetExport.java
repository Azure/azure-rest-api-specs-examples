/** Samples for Jobs GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/JobsGetExport.json
     */
    /**
     * Sample code: JobsGetExport.
     *
     * @param manager Entry point to DataBoxManager.
     */
    public static void jobsGetExport(com.azure.resourcemanager.databox.DataBoxManager manager) {
        manager
            .jobs()
            .getByResourceGroupWithResponse(
                "YourResourceGroupName", "TestJobName1", "details", com.azure.core.util.Context.NONE);
    }
}
