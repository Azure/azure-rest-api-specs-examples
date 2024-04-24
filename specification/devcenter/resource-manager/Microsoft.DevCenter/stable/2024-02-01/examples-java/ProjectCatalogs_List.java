
/**
 * Samples for ProjectCatalogs List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/ProjectCatalogs_List.json
     */
    /**
     * Sample code: ProjectCatalogs_List.
     * 
     * @param manager Entry point to DevCenterManager.
     */
    public static void projectCatalogsList(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.projectCatalogs().list("rg1", "DevProject", null, com.azure.core.util.Context.NONE);
    }
}
