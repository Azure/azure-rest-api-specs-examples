
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
     * ApiManagementCreateODataApi.json
     */
    /**
     * Sample code: ApiManagementCreateODataApi.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementCreateODataApi(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apis().define("tempgroup").withExistingService("rg1", "apimService1")
            .withValue("https://services.odata.org/TripPinWebApiService/$metadata").withFormat(ContentFormat.ODATA_LINK)
            .withDisplayName("apiname1463").withServiceUrl("https://services.odata.org/TripPinWebApiService")
            .withPath("odata-api").withProtocols(Arrays.asList(Protocol.HTTP, Protocol.HTTPS))
            .withDescription("apidescription5200").withApiType(ApiType.ODATA).create();
    }
}
