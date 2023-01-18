/** Samples for ManagementAssociations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationsmanagement/resource-manager/Microsoft.OperationsManagement/preview/2015-11-01-preview/examples/ManagementAssociationDelete.json
     */
    /**
     * Sample code: SolutionDelete.
     *
     * @param manager Entry point to OperationsManagementManager.
     */
    public static void solutionDelete(
        com.azure.resourcemanager.operationsmanagement.OperationsManagementManager manager) {
        manager
            .managementAssociations()
            .deleteWithResponse(
                "rg1",
                "providerName",
                "resourceType",
                "resourceName",
                "managementAssociationName",
                com.azure.core.util.Context.NONE);
    }
}
