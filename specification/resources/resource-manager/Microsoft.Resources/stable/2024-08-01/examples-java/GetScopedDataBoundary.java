
import com.azure.resourcemanager.resources.models.DefaultName;

/**
 * Samples for DataBoundaries GetScope.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Resources/stable/2024-08-01/examples/GetScopedDataBoundary.
     * json
     */
    /**
     * Sample code: Get data boundary at scope.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDataBoundaryAtScope(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().dataBoundaryClient().getDataBoundaries().getScopeWithResponse(
            "subscriptions/11111111-1111-1111-1111-111111111111/resourcegroups/my-resource-group", DefaultName.DEFAULT,
            com.azure.core.util.Context.NONE);
    }
}
