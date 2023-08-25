/** Samples for NamedValue Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetNamedValue.json
     */
    /**
     * Sample code: ApiManagementGetNamedValue.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetNamedValue(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .namedValues()
            .getWithResponse("rg1", "apimService1", "testarmTemplateproperties2", com.azure.core.util.Context.NONE);
    }
}
