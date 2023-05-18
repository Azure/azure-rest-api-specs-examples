/** Samples for EnvironmentTypes Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/EnvironmentTypes_Get.json
     */
    /**
     * Sample code: EnvironmentTypes_Get.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void environmentTypesGet(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.environmentTypes().getWithResponse("rg1", "Contoso", "DevTest", com.azure.core.util.Context.NONE);
    }
}
