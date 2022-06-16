import com.azure.core.util.Context;

/** Samples for ApiDiagnostic GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadApiDiagnostic.json
     */
    /**
     * Sample code: ApiManagementHeadApiDiagnostic.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadApiDiagnostic(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiDiagnostics()
            .getEntityTagWithResponse(
                "rg1", "apimService1", "57d1f7558aa04f15146d9d8a", "applicationinsights", Context.NONE);
    }
}
