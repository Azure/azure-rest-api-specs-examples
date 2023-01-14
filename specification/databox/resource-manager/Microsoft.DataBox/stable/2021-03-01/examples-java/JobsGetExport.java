/** Samples for Jobs GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2021-03-01/examples/JobsGetExport.json
     */
    /**
     * Sample code: JobsGetExport.
     *
     * @param manager Entry point to DataBoxManager.
     */
    public static void jobsGetExport(com.azure.resourcemanager.databox.DataBoxManager manager) {
        manager
            .jobs()
            .getByResourceGroupWithResponse("SdkRg8091", "SdkJob6429", "details", com.azure.core.util.Context.NONE);
    }
}
