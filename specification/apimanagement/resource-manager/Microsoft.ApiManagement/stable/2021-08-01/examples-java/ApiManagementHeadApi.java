import com.azure.core.util.Context;

/** Samples for Api GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadApi.json
     */
    /**
     * Sample code: ApiManagementHeadApi.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadApi(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apis().getEntityTagWithResponse("rg1", "apimService1", "57d1f7558aa04f15146d9d8a", Context.NONE);
    }
}
