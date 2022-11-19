import com.azure.core.util.Context;

/** Samples for EnvironmentTypes Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-10-12-preview/examples/EnvironmentTypes_Get.json
     */
    /**
     * Sample code: EnvironmentTypes_Get.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void environmentTypesGet(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.environmentTypes().getWithResponse("rg1", "Contoso", "{environmentTypeName}", Context.NONE);
    }
}
