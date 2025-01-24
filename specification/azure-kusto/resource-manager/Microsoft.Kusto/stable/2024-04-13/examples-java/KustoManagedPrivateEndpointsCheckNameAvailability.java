
import com.azure.resourcemanager.kusto.models.ManagedPrivateEndpointsCheckNameRequest;

/**
 * Samples for ManagedPrivateEndpoints CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/
     * KustoManagedPrivateEndpointsCheckNameAvailability.json
     */
    /**
     * Sample code: KustoManagedPrivateEndpointsCheckNameAvailability.
     * 
     * @param manager Entry point to KustoManager.
     */
    public static void
        kustoManagedPrivateEndpointsCheckNameAvailability(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.managedPrivateEndpoints().checkNameAvailabilityWithResponse("kustorptest", "kustoCluster",
            new ManagedPrivateEndpointsCheckNameRequest().withName("pme1"), com.azure.core.util.Context.NONE);
    }
}
