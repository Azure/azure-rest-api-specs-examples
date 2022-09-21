import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-07-07/examples/KustoPrivateEndpointConnectionsList.json
     */
    /**
     * Sample code: KustoPrivateEndpointConnectionsList.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoPrivateEndpointConnectionsList(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.privateEndpointConnections().list("kustorptest", "kustoCluster", Context.NONE);
    }
}
