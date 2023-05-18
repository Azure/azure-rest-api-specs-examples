/** Samples for Projects List. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/Projects_ListBySubscription.json
     */
    /**
     * Sample code: Projects_ListBySubscription.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void projectsListBySubscription(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.projects().list(null, com.azure.core.util.Context.NONE);
    }
}
