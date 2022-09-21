import com.azure.core.util.Context;

/** Samples for ManagedPrivateEndpoints Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-07-07/examples/KustoManagedPrivateEndpointsGet.json
     */
    /**
     * Sample code: KustoManagedPrivateEndpointsGet.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoManagedPrivateEndpointsGet(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .managedPrivateEndpoints()
            .getWithResponse("kustorptest", "kustoCluster", "managedPrivateEndpointTest", Context.NONE);
    }
}
