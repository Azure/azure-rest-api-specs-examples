import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoPrivateEndpointConnectionsList.json
     */
    /**
     * Sample code: KustoPrivateEndpointConnectionsList.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoPrivateEndpointConnectionsList(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.privateEndpointConnections().list("kustorptest", "kustoclusterrptest4", Context.NONE);
    }
}
