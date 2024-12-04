
/**
 * Samples for EnvironmentDefinitions ListByProjectCatalog.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/
     * EnvironmentDefinitions_ListByProjectCatalog.json
     */
    /**
     * Sample code: EnvironmentDefinitions_ListByProjectCatalog.
     * 
     * @param manager Entry point to DevCenterManager.
     */
    public static void
        environmentDefinitionsListByProjectCatalog(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.environmentDefinitions().listByProjectCatalog("rg1", "DevProject", "myCatalog",
            com.azure.core.util.Context.NONE);
    }
}
