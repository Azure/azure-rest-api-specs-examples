/** Samples for ProjectEnvironmentTypes Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/ProjectEnvironmentTypes_Get.json
     */
    /**
     * Sample code: ProjectEnvironmentTypes_Get.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void projectEnvironmentTypesGet(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager
            .projectEnvironmentTypes()
            .getWithResponse("rg1", "ContosoProj", "DevTest", com.azure.core.util.Context.NONE);
    }
}
