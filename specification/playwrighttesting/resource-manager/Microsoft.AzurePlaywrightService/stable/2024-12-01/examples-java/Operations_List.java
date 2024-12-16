
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-12-01/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     * 
     * @param manager Entry point to PlaywrightTestingManager.
     */
    public static void operationsList(com.azure.resourcemanager.playwrighttesting.PlaywrightTestingManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
