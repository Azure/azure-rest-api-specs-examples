```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.ExportApi;
import com.azure.resourcemanager.apimanagement.models.ExportFormat;

/** Samples for ApiExport Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetApiExportInOpenApi3dot0.json
     */
    /**
     * Sample code: ApiManagementGetApiExportInOpenApi3dot0.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetApiExportInOpenApi3dot0(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiExports()
            .getWithResponse("rg1", "apimService1", "aid9676", ExportFormat.OPENAPI_LINK, ExportApi.TRUE, Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
