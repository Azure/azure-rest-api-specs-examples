
/**
 * Samples for Linker Validate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/
     * ValidateLinkerSuccess.json
     */
    /**
     * Sample code: ValidateLinkerSuccess.
     * 
     * @param manager Entry point to ServiceLinkerManager.
     */
    public static void validateLinkerSuccess(com.azure.resourcemanager.servicelinker.ServiceLinkerManager manager) {
        manager.linkers().validate(
            "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app",
            "linkName", com.azure.core.util.Context.NONE);
    }
}
