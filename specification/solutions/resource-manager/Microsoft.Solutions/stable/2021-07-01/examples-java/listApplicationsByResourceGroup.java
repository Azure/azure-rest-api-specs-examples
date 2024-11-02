
/**
 * Samples for Applications ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/
     * listApplicationsByResourceGroup.json
     */
    /**
     * Sample code: Lists all the applications within a resource group.
     * 
     * @param manager Entry point to ApplicationManager.
     */
    public static void listsAllTheApplicationsWithinAResourceGroup(
        com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        manager.applications().listByResourceGroup("rg", com.azure.core.util.Context.NONE);
    }
}
