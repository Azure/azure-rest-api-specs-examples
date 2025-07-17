
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01-preview/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     * 
     * @param manager Entry point to PlaywrightManager.
     */
    public static void operationsList(com.azure.resourcemanager.playwright.PlaywrightManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
