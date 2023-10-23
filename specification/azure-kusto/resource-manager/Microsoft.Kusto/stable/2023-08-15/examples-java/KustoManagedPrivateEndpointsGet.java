/** Samples for ManagedPrivateEndpoints Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoManagedPrivateEndpointsGet.json
     */
    /**
     * Sample code: KustoManagedPrivateEndpointsGet.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoManagedPrivateEndpointsGet(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .managedPrivateEndpoints()
            .getWithResponse(
                "kustorptest", "kustoCluster", "managedPrivateEndpointTest", com.azure.core.util.Context.NONE);
    }
}
