import com.azure.core.util.Context;

/** Samples for ManagedPrivateEndpoints Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-07-07/examples/KustoManagedPrivateEndpointsDelete.json
     */
    /**
     * Sample code: ManagedPrivateEndpointsDelete.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void managedPrivateEndpointsDelete(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .managedPrivateEndpoints()
            .delete("kustorptest", "kustoCluster", "managedPrivateEndpointTest", Context.NONE);
    }
}
