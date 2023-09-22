/** Samples for Applications Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/deleteApplication.json
     */
    /**
     * Sample code: Delete managed application.
     *
     * @param manager Entry point to ApplicationManager.
     */
    public static void deleteManagedApplication(
        com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        manager.applications().delete("rg", "myManagedApplication", com.azure.core.util.Context.NONE);
    }
}
