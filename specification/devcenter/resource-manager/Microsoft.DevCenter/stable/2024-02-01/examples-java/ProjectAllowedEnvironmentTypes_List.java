
/**
 * Samples for ProjectAllowedEnvironmentTypes List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/
     * ProjectAllowedEnvironmentTypes_List.json
     */
    /**
     * Sample code: ProjectAllowedEnvironmentTypes_List.
     * 
     * @param manager Entry point to DevCenterManager.
     */
    public static void
        projectAllowedEnvironmentTypesList(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.projectAllowedEnvironmentTypes().list("rg1", "Contoso", null, com.azure.core.util.Context.NONE);
    }
}
