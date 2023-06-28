/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/ListOperations.json
     */
    /**
     * Sample code: ListOperations.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void listOperations(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
