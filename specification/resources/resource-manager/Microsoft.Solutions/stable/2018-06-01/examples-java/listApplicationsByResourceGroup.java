/** Samples for Applications ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/listApplicationsByResourceGroup.json
     */
    /**
     * Sample code: Lists applications.
     *
     * @param manager Entry point to ApplicationManager.
     */
    public static void listsApplications(com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        manager.applications().listByResourceGroup("rg", com.azure.core.util.Context.NONE);
    }
}
