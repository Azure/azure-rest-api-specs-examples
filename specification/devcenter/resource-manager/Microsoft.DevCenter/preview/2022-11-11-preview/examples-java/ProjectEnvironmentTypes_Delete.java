import com.azure.core.util.Context;

/** Samples for ProjectEnvironmentTypes Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/ProjectEnvironmentTypes_Delete.json
     */
    /**
     * Sample code: ProjectEnvironmentTypes_Delete.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void projectEnvironmentTypesDelete(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.projectEnvironmentTypes().deleteWithResponse("rg1", "ContosoProj", "DevTest", Context.NONE);
    }
}
