/** Samples for NamedValue ListValue. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementNamedValueListValue.json
     */
    /**
     * Sample code: ApiManagementNamedValueListValue.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementNamedValueListValue(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .namedValues()
            .listValueWithResponse(
                "rg1", "apimService1", "testarmTemplateproperties2", com.azure.core.util.Context.NONE);
    }
}
