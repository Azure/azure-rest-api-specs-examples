import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2022-10-31/examples/DigitalTwinsOperationsList_example.json
     */
    /**
     * Sample code: Get available operations.
     *
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void getAvailableOperations(com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager.operations().list(Context.NONE);
    }
}
