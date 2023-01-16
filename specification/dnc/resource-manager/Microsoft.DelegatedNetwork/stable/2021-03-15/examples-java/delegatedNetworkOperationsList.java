/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnc/resource-manager/Microsoft.DelegatedNetwork/stable/2021-03-15/examples/delegatedNetworkOperationsList.json
     */
    /**
     * Sample code: Get available operations.
     *
     * @param manager Entry point to DelegatedNetworkManager.
     */
    public static void getAvailableOperations(
        com.azure.resourcemanager.delegatednetwork.DelegatedNetworkManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
