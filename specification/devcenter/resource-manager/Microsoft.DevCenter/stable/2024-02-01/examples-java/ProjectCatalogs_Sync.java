
/**
 * Samples for ProjectCatalogs Sync.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/ProjectCatalogs_Sync.json
     */
    /**
     * Sample code: ProjectCatalogs_Sync.
     * 
     * @param manager Entry point to DevCenterManager.
     */
    public static void projectCatalogsSync(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.projectCatalogs().sync("rg1", "DevProject", "CentralCatalog", com.azure.core.util.Context.NONE);
    }
}
