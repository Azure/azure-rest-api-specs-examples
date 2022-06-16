import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/stable/2022-05-01/examples/OperationsList.json
     */
    /**
     * Sample code: GetConfiguration.
     *
     * @param manager Entry point to ServiceLinkerManager.
     */
    public static void getConfiguration(com.azure.resourcemanager.servicelinker.ServiceLinkerManager manager) {
        manager.operations().list(Context.NONE);
    }
}
