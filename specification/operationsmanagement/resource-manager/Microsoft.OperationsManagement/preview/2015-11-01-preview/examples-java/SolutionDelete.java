/** Samples for Solutions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationsmanagement/resource-manager/Microsoft.OperationsManagement/preview/2015-11-01-preview/examples/SolutionDelete.json
     */
    /**
     * Sample code: SolutionDelete.
     *
     * @param manager Entry point to OperationsManagementManager.
     */
    public static void solutionDelete(
        com.azure.resourcemanager.operationsmanagement.OperationsManagementManager manager) {
        manager.solutions().delete("rg1", "solution1", com.azure.core.util.Context.NONE);
    }
}
