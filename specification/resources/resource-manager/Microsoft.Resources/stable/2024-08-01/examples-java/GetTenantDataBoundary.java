
import com.azure.resourcemanager.resources.models.DefaultName;

/**
 * Samples for DataBoundaries GetTenant.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Resources/stable/2024-08-01/examples/GetTenantDataBoundary.
     * json
     */
    /**
     * Sample code: Get data boundary for tenant.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDataBoundaryForTenant(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().dataBoundaryClient().getDataBoundaries()
            .getTenantWithResponse(DefaultName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
