import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-05-01-preview/examples/ListOperations.json
     */
    /**
     * Sample code: ListOperations.
     *
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void listOperations(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager.operations().list(Context.NONE);
    }
}
