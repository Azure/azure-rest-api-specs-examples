import com.azure.resourcemanager.apimanagement.models.ExportApi;
import com.azure.resourcemanager.apimanagement.models.ExportFormat;

/** Samples for ApiExport Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetApiExportInOpenApi2dot0.json
     */
    /**
     * Sample code: ApiManagementGetApiExportInOpenApi2dot0.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetApiExportInOpenApi2dot0(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiExports()
            .getWithResponse(
                "rg1",
                "apimService1",
                "echo-api",
                ExportFormat.SWAGGER_LINK,
                ExportApi.TRUE,
                com.azure.core.util.Context.NONE);
    }
}
