
/**
 * Samples for ManagedPrivateEndpoints List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/
     * KustoManagedPrivateEndpointsList.json
     */
    /**
     * Sample code: KustoManagedPrivateEndpointsList.
     * 
     * @param manager Entry point to KustoManager.
     */
    public static void kustoManagedPrivateEndpointsList(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.managedPrivateEndpoints().list("kustorptest", "kustoCluster", com.azure.core.util.Context.NONE);
    }
}
