import com.azure.core.util.Context;

/** Samples for ApiPortals List. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/ApiPortals_List.json
     */
    /**
     * Sample code: ApiPortals_List.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void apiPortalsList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getApiPortals()
            .list("myResourceGroup", "myservice", Context.NONE);
    }
}
