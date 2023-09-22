/** Samples for Applications List. */
public final class Main {
    /*
     * x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/listApplicationsByResourceGroup.json
     */
    /**
     * Sample code: Lists all the applications within a subscription.
     *
     * @param manager Entry point to ApplicationManager.
     */
    public static void listsAllTheApplicationsWithinASubscription(
        com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        manager.applications().list(com.azure.core.util.Context.NONE);
    }
}
