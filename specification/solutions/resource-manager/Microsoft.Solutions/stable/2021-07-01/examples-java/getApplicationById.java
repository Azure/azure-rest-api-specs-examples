/** Samples for Applications GetById. */
public final class Main {
    /*
     * x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/getApplicationById.json
     */
    /**
     * Sample code: Gets the managed application.
     *
     * @param manager Entry point to ApplicationManager.
     */
    public static void getsTheManagedApplication(
        com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        manager
            .applications()
            .getByIdWithResponse(
                "subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applications/myManagedApplication",
                com.azure.core.util.Context.NONE);
    }
}
