import com.azure.core.util.Context;

/** Samples for PrivateLinkResources Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoPrivateLinkResourcesGet.json
     */
    /**
     * Sample code: Gets private endpoint connection.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void getsPrivateEndpointConnection(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.privateLinkResources().getWithResponse("kustorptest", "kustoclusterrptest4", "cluster", Context.NONE);
    }
}
