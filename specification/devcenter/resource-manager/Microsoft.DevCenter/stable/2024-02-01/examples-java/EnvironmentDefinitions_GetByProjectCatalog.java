
/**
 * Samples for EnvironmentDefinitions GetByProjectCatalog.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/
     * EnvironmentDefinitions_GetByProjectCatalog.json
     */
    /**
     * Sample code: EnvironmentDefinitions_GetByProjectCatalog.
     * 
     * @param manager Entry point to DevCenterManager.
     */
    public static void
        environmentDefinitionsGetByProjectCatalog(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.environmentDefinitions().getByProjectCatalogWithResponse("rg1", "DevProject", "myCatalog",
            "myEnvironmentDefinition", com.azure.core.util.Context.NONE);
    }
}
