import com.azure.core.util.Context;
import com.azure.resourcemanager.kusto.models.ManagedPrivateEndpoint;

/** Samples for ManagedPrivateEndpoints Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoManagedPrivateEndpointsUpdate.json
     */
    /**
     * Sample code: KustoManagedPrivateEndpointsUpdate.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoManagedPrivateEndpointsUpdate(com.azure.resourcemanager.kusto.KustoManager manager) {
        ManagedPrivateEndpoint resource =
            manager
                .managedPrivateEndpoints()
                .getWithResponse("kustorptest", "kustoclusterrptest4", "kustoManagedPrivateEndpoint1", Context.NONE)
                .getValue();
        resource
            .update()
            .withPrivateLinkResourceId(
                "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Storage/storageAccounts/storageAccountTest")
            .withGroupId("blob")
            .withRequestMessage("Please Approve Managed Private Endpoint Request.")
            .apply();
    }
}
