import com.azure.core.util.Context;

/** Samples for Tenants List. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-01-01/examples/GetTenants.json
     */
    /**
     * Sample code: GetAllTenants.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAllTenants(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().subscriptionClient().getTenants().list(Context.NONE);
    }
}
