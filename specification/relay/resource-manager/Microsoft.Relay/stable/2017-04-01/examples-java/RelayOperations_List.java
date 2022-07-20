import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/RelayOperations_List.json
     */
    /**
     * Sample code: RelayOperationsList.
     *
     * @param manager Entry point to RelayManager.
     */
    public static void relayOperationsList(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.operations().list(Context.NONE);
    }
}
