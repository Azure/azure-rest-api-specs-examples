
/**
 * Samples for ProjectCatalogs GetSyncErrorDetails.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/
     * ProjectCatalogs_GetSyncErrorDetails.json
     */
    /**
     * Sample code: ProjectCatalogs_GetSyncErrorDetails.
     * 
     * @param manager Entry point to DevCenterManager.
     */
    public static void
        projectCatalogsGetSyncErrorDetails(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.projectCatalogs().getSyncErrorDetailsWithResponse("rg1", "DevProject", "CentralCatalog",
            com.azure.core.util.Context.NONE);
    }
}
