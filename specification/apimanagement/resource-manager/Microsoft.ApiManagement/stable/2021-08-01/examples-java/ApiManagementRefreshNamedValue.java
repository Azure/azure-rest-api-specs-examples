import com.azure.core.util.Context;

/** Samples for NamedValue RefreshSecret. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementRefreshNamedValue.json
     */
    /**
     * Sample code: ApiManagementRefreshNamedValue.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementRefreshNamedValue(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.namedValues().refreshSecret("rg1", "apimService1", "testprop2", Context.NONE);
    }
}
