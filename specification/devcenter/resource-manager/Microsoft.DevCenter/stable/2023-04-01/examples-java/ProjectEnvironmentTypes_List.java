/** Samples for ProjectEnvironmentTypes List. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/ProjectEnvironmentTypes_List.json
     */
    /**
     * Sample code: ProjectEnvironmentTypes_List.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void projectEnvironmentTypesList(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.projectEnvironmentTypes().list("rg1", "ContosoProj", null, com.azure.core.util.Context.NONE);
    }
}
