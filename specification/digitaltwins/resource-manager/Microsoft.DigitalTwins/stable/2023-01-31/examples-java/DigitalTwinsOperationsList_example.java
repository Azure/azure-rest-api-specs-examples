/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2023-01-31/examples/DigitalTwinsOperationsList_example.json
     */
    /**
     * Sample code: Get available operations.
     *
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void getAvailableOperations(com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
