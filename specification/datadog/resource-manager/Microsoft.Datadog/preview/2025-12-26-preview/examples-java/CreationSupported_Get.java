
/**
 * Samples for CreationSupported Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-26-preview/CreationSupported_Get.json
     */
    /**
     * Sample code: CreationSupported_Get.
     * 
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void creationSupportedGet(com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager.creationSupporteds().getWithResponse("00000000-0000-0000-0000", com.azure.core.util.Context.NONE);
    }
}
