/** Samples for Tag DetachFromApi. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteApiTag.json
     */
    /**
     * Sample code: ApiManagementDeleteApiTag.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteApiTag(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .tags()
            .detachFromApiWithResponse(
                "rg1",
                "apimService1",
                "59d5b28d1f7fab116c282650",
                "59d5b28e1f7fab116402044e",
                com.azure.core.util.Context.NONE);
    }
}
