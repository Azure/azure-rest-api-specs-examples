
/**
 * Samples for ProjectCatalogs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/ProjectCatalogs_Delete.
     * json
     */
    /**
     * Sample code: ProjectCatalogs_Delete.
     * 
     * @param manager Entry point to DevCenterManager.
     */
    public static void projectCatalogsDelete(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.projectCatalogs().delete("rg1", "DevProject", "CentralCatalog", com.azure.core.util.Context.NONE);
    }
}
