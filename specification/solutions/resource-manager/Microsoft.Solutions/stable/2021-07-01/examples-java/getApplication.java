/** Samples for Applications GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/getApplication.json
     */
    /**
     * Sample code: Get a managed application.
     *
     * @param manager Entry point to ApplicationManager.
     */
    public static void getAManagedApplication(
        com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        manager
            .applications()
            .getByResourceGroupWithResponse("rg", "myManagedApplication", com.azure.core.util.Context.NONE);
    }
}
