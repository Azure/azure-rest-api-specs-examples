/** Samples for DevBoxDefinitions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/DevBoxDefinitions_Get.json
     */
    /**
     * Sample code: DevBoxDefinitions_Get.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void devBoxDefinitionsGet(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.devBoxDefinitions().getWithResponse("rg1", "Contoso", "WebDevBox", com.azure.core.util.Context.NONE);
    }
}
