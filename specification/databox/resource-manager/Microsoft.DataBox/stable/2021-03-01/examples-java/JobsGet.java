/** Samples for Jobs GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2021-03-01/examples/JobsGet.json
     */
    /**
     * Sample code: JobsGet.
     *
     * @param manager Entry point to DataBoxManager.
     */
    public static void jobsGet(com.azure.resourcemanager.databox.DataBoxManager manager) {
        manager
            .jobs()
            .getByResourceGroupWithResponse("SdkRg5154", "SdkJob952", "details", com.azure.core.util.Context.NONE);
    }
}
