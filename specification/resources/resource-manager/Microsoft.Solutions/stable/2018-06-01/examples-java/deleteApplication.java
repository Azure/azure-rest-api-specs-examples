/** Samples for Applications Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/deleteApplication.json
     */
    /**
     * Sample code: Deletes a managed application.
     *
     * @param manager Entry point to ApplicationManager.
     */
    public static void deletesAManagedApplication(
        com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        manager.applications().delete("rg", "myManagedApplication", com.azure.core.util.Context.NONE);
    }
}
