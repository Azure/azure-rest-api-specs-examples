/** Samples for Applications ListAllowedUpgradePlans. */
public final class Main {
    /*
     * x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/listAllowedUpgradePlans.json
     */
    /**
     * Sample code: List allowed upgrade plans for application.
     *
     * @param manager Entry point to ApplicationManager.
     */
    public static void listAllowedUpgradePlansForApplication(
        com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        manager
            .applications()
            .listAllowedUpgradePlansWithResponse("rg", "myManagedApplication", com.azure.core.util.Context.NONE);
    }
}
