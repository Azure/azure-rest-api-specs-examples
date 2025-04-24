
import com.azure.resourcemanager.apimanagement.models.ApiType;
import com.azure.resourcemanager.apimanagement.models.ContentFormat;
import com.azure.resourcemanager.apimanagement.models.Protocol;
import java.util.Arrays;

/**
 * Samples for Api CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateGrpcApi.json
     */
    /**
     * Sample code: ApiManagementCreateGrpcApi.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementCreateGrpcApi(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apis().define("tempgroup").withExistingService("rg1", "apimService1")
            .withValue(
                "https://raw.githubusercontent.com/kedacore/keda/main/pkg/scalers/externalscaler/externalscaler.proto")
            .withFormat(ContentFormat.GRPC_LINK).withDisplayName("apiname1463")
            .withServiceUrl("https://your-api-hostname/samples").withPath("grpc-api")
            .withProtocols(Arrays.asList(Protocol.HTTPS)).withDescription("apidescription5200")
            .withApiType(ApiType.GRPC).create();
    }
}
