/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ListOperations.json
     */
    /**
     * Sample code: ListOperations.
     *
     * @param manager Entry point to ServiceFabricManager.
     */
    public static void listOperations(com.azure.resourcemanager.servicefabric.ServiceFabricManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
