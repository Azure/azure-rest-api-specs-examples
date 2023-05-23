/** Samples for Jobs ListCredentials. */
public final class Main {
    /*
     * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/JobsListCredentials.json
     */
    /**
     * Sample code: JobsListCredentials.
     *
     * @param manager Entry point to DataBoxManager.
     */
    public static void jobsListCredentials(com.azure.resourcemanager.databox.DataBoxManager manager) {
        manager.jobs().listCredentials("YourResourceGroupName", "TestJobName1", com.azure.core.util.Context.NONE);
    }
}
