/** Samples for Jobs GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2021-03-01/examples/JobsGetCmk.json
     */
    /**
     * Sample code: JobsGetCmk.
     *
     * @param manager Entry point to DataBoxManager.
     */
    public static void jobsGetCmk(com.azure.resourcemanager.databox.DataBoxManager manager) {
        manager
            .jobs()
            .getByResourceGroupWithResponse("SdkRg7937", "SdkJob1735", "details", com.azure.core.util.Context.NONE);
    }
}
