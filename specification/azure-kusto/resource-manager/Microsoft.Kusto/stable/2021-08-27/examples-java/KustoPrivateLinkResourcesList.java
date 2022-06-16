import com.azure.core.util.Context;

/** Samples for PrivateLinkResources List. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoPrivateLinkResourcesList.json
     */
    /**
     * Sample code: Gets private endpoint connections.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void getsPrivateEndpointConnections(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.privateLinkResources().list("kustorptest", "kustoclusterrptest4", Context.NONE);
    }
}
