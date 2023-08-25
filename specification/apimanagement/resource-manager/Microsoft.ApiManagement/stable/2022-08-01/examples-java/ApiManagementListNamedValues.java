/** Samples for NamedValue ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListNamedValues.json
     */
    /**
     * Sample code: ApiManagementListNamedValues.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListNamedValues(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .namedValues()
            .listByService("rg1", "apimService1", null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
