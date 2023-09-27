/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/playwrighttesting/resource-manager/Microsoft.AzurePlaywrightService/preview/2023-10-01-preview/examples/Operations_List.json
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
