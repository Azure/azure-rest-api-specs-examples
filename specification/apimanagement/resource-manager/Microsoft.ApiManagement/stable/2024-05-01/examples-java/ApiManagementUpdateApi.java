
import com.azure.resourcemanager.apimanagement.models.ApiContract;

/**
 * Samples for Api Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementUpdateApi.json
     */
    /**
     * Sample code: ApiManagementUpdateApi.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateApi(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        ApiContract resource = manager.apis()
            .getWithResponse("rg1", "apimService1", "echo-api", com.azure.core.util.Context.NONE).getValue();
        resource.update().withDisplayName("Echo API New").withServiceUrl("http://echoapi.cloudapp.net/api2")
            .withPath("newecho").withIfMatch("*").apply();
    }
}
