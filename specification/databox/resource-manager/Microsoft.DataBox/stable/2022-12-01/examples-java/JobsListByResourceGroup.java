/** Samples for Jobs ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/JobsListByResourceGroup.json
     */
    /**
     * Sample code: JobsListByResourceGroup.
     *
     * @param manager Entry point to DataBoxManager.
     */
    public static void jobsListByResourceGroup(com.azure.resourcemanager.databox.DataBoxManager manager) {
        manager.jobs().listByResourceGroup("YourResourceGroupName", null, com.azure.core.util.Context.NONE);
    }
}
