/** Samples for NamedValue GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadNamedValue.json
     */
    /**
     * Sample code: ApiManagementHeadNamedValue.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadNamedValue(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .namedValues()
            .getEntityTagWithResponse(
                "rg1", "apimService1", "testarmTemplateproperties2", com.azure.core.util.Context.NONE);
    }
}
