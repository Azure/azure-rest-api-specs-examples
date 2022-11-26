import com.azure.core.util.Context;

/** Samples for ProjectAllowedEnvironmentTypes Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/ProjectAllowedEnvironmentTypes_Get.json
     */
    /**
     * Sample code: ProjectAllowedEnvironmentTypes_Get.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void projectAllowedEnvironmentTypesGet(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.projectAllowedEnvironmentTypes().getWithResponse("rg1", "Contoso", "DevTest", Context.NONE);
    }
}
