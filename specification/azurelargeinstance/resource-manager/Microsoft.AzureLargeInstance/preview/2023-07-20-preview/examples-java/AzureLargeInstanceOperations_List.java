
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurelargeinstance/resource-manager/Microsoft.AzureLargeInstance/preview/2023-07-20-preview/
     * examples/AzureLargeInstanceOperations_List.json
     */
    /**
     * Sample code: Operations_List.
     * 
     * @param manager Entry point to LargeInstanceManager.
     */
    public static void operationsList(com.azure.resourcemanager.largeinstance.LargeInstanceManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
