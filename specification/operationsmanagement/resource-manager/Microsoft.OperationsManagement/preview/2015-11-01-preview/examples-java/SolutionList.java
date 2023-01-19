/** Samples for Solutions ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationsmanagement/resource-manager/Microsoft.OperationsManagement/preview/2015-11-01-preview/examples/SolutionList.json
     */
    /**
     * Sample code: SolutionList.
     *
     * @param manager Entry point to OperationsManagementManager.
     */
    public static void solutionList(
        com.azure.resourcemanager.operationsmanagement.OperationsManagementManager manager) {
        manager.solutions().listByResourceGroupWithResponse("rg1", com.azure.core.util.Context.NONE);
    }
}
