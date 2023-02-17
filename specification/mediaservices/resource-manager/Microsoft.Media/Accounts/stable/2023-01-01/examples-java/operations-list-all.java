/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Accounts/stable/2023-01-01/examples/operations-list-all.json
     */
    /**
     * Sample code: List Operations.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listOperations(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.operations().listWithResponse(com.azure.core.util.Context.NONE);
    }
}
