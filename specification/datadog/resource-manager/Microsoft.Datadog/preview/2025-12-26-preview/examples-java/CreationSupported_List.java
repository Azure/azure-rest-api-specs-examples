
/**
 * Samples for CreationSupported List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-26-preview/CreationSupported_List.json
     */
    /**
     * Sample code: CreationSupported_List.
     * 
     * @param manager Entry point to MicrosoftDatadogManager.
     */
    public static void creationSupportedList(com.azure.resourcemanager.datadog.MicrosoftDatadogManager manager) {
        manager.creationSupporteds().list("00000000-0000-0000-0000", com.azure.core.util.Context.NONE);
    }
}
