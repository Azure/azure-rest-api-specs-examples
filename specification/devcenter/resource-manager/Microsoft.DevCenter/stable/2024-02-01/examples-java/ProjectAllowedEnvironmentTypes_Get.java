
/**
 * Samples for ProjectAllowedEnvironmentTypes Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/
     * ProjectAllowedEnvironmentTypes_Get.json
     */
    /**
     * Sample code: ProjectAllowedEnvironmentTypes_Get.
     * 
     * @param manager Entry point to DevCenterManager.
     */
    public static void projectAllowedEnvironmentTypesGet(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.projectAllowedEnvironmentTypes().getWithResponse("rg1", "Contoso", "DevTest",
            com.azure.core.util.Context.NONE);
    }
}
