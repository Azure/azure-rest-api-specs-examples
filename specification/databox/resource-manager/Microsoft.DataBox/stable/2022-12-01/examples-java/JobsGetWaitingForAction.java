/** Samples for Jobs GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/JobsGetWaitingForAction.json
     */
    /**
     * Sample code: JobsGetWaitingForAction.
     *
     * @param manager Entry point to DataBoxManager.
     */
    public static void jobsGetWaitingForAction(com.azure.resourcemanager.databox.DataBoxManager manager) {
        manager
            .jobs()
            .getByResourceGroupWithResponse(
                "YourResourceGroupName", "TestJobName1", "details", com.azure.core.util.Context.NONE);
    }
}
