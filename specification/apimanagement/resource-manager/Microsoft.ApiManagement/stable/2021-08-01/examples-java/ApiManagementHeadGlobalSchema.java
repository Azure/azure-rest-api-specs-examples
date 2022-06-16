import com.azure.core.util.Context;

/** Samples for GlobalSchema GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadGlobalSchema.json
     */
    /**
     * Sample code: ApiManagementHeadApi.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadApi(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.globalSchemas().getEntityTagWithResponse("rg1", "apimService1", "myschema", Context.NONE);
    }
}
