/** Samples for ResourceProvider ListOperations. */
public final class Main {
    /*
     * x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/listSolutionsOperations.json
     */
    /**
     * Sample code: List Solutions operations.
     *
     * @param manager Entry point to ApplicationManager.
     */
    public static void listSolutionsOperations(
        com.azure.resourcemanager.managedapplications.ApplicationManager manager) {
        manager.resourceProviders().listOperations(com.azure.core.util.Context.NONE);
    }
}
