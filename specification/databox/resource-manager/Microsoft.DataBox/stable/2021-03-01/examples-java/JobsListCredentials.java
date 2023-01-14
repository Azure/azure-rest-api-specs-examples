/** Samples for Jobs ListCredentials. */
public final class Main {
    /*
     * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2021-03-01/examples/JobsListCredentials.json
     */
    /**
     * Sample code: JobsListCredentials.
     *
     * @param manager Entry point to DataBoxManager.
     */
    public static void jobsListCredentials(com.azure.resourcemanager.databox.DataBoxManager manager) {
        manager.jobs().listCredentials("bvttoolrg6", "TJ-636646322037905056", com.azure.core.util.Context.NONE);
    }
}
