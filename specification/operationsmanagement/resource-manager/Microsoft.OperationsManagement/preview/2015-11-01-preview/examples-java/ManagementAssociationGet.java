/** Samples for ManagementAssociations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/operationsmanagement/resource-manager/Microsoft.OperationsManagement/preview/2015-11-01-preview/examples/ManagementAssociationGet.json
     */
    /**
     * Sample code: SolutionGet.
     *
     * @param manager Entry point to OperationsManagementManager.
     */
    public static void solutionGet(com.azure.resourcemanager.operationsmanagement.OperationsManagementManager manager) {
        manager
            .managementAssociations()
            .getWithResponse(
                "rg1",
                "providerName",
                "resourceType",
                "resourceName",
                "managementAssociation1",
                com.azure.core.util.Context.NONE);
    }
}
