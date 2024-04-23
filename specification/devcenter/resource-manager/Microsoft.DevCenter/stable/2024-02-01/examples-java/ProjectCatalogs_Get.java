
/**
 * Samples for ProjectCatalogs Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/ProjectCatalogs_Get.json
     */
    /**
     * Sample code: ProjectCatalogs_Get.
     * 
     * @param manager Entry point to DevCenterManager.
     */
    public static void projectCatalogsGet(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.projectCatalogs().getWithResponse("rg1", "DevProject", "CentralCatalog",
            com.azure.core.util.Context.NONE);
    }
}
