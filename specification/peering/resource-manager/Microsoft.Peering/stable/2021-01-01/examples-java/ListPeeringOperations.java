/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-01-01/examples/ListPeeringOperations.json
     */
    /**
     * Sample code: List peering operations.
     *
     * @param manager Entry point to PeeringManager.
     */
    public static void listPeeringOperations(com.azure.resourcemanager.peering.PeeringManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
