
/**
 * Samples for Linker Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/Linker.
     * json
     */
    /**
     * Sample code: Linker.
     * 
     * @param manager Entry point to ServiceLinkerManager.
     */
    public static void linker(com.azure.resourcemanager.servicelinker.ServiceLinkerManager manager) {
        manager.linkers().getWithResponse(
            "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app",
            "linkName", com.azure.core.util.Context.NONE);
    }
}
