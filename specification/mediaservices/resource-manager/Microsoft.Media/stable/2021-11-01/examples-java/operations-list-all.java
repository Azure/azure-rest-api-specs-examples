import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/operations-list-all.json
     */
    /**
     * Sample code: List Operations.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listOperations(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.operations().listWithResponse(Context.NONE);
    }
}
