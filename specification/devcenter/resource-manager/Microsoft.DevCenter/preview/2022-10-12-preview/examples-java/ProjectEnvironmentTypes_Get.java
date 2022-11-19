import com.azure.core.util.Context;

/** Samples for ProjectEnvironmentTypes Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-10-12-preview/examples/ProjectEnvironmentTypes_Get.json
     */
    /**
     * Sample code: ProjectEnvironmentTypes_Get.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void projectEnvironmentTypesGet(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.projectEnvironmentTypes().getWithResponse("rg1", "ContosoProj", "{environmentTypeName}", Context.NONE);
    }
}
