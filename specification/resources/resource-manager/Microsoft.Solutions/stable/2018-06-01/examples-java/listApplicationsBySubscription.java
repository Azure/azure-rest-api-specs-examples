/** Samples for Applications List. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/listApplicationsBySubscription.json
     */
    /**
     * Sample code: Lists applications by subscription.
     *
     * @param manager Entry point to ApplicationManager.
     */
    public static void listsApplicationsBySubscription(
        com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        manager.applications().list(com.azure.core.util.Context.NONE);
    }
}
