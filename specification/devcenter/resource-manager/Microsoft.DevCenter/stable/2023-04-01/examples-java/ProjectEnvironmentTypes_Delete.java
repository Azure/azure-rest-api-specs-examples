/** Samples for ProjectEnvironmentTypes Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/ProjectEnvironmentTypes_Delete.json
     */
    /**
     * Sample code: ProjectEnvironmentTypes_Delete.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void projectEnvironmentTypesDelete(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager
            .projectEnvironmentTypes()
            .deleteWithResponse("rg1", "ContosoProj", "DevTest", com.azure.core.util.Context.NONE);
    }
}
