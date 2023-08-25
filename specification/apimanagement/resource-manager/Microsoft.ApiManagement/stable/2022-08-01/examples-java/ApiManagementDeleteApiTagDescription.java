/** Samples for ApiTagDescription Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteApiTagDescription.json
     */
    /**
     * Sample code: ApiManagementDeleteApiTagDescription.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteApiTagDescription(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiTagDescriptions()
            .deleteWithResponse(
                "rg1",
                "apimService1",
                "59d5b28d1f7fab116c282650",
                "59d5b28e1f7fab116402044e",
                "*",
                com.azure.core.util.Context.NONE);
    }
}
