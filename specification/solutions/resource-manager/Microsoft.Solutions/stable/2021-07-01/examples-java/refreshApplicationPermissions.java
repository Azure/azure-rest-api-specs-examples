/** Samples for Applications RefreshPermissions. */
public final class Main {
    /*
     * x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/refreshApplicationPermissions.json
     */
    /**
     * Sample code: Refresh managed application permissions.
     *
     * @param manager Entry point to ApplicationManager.
     */
    public static void refreshManagedApplicationPermissions(
        com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        manager.applications().refreshPermissions("rg", "myManagedApplication", com.azure.core.util.Context.NONE);
    }
}
