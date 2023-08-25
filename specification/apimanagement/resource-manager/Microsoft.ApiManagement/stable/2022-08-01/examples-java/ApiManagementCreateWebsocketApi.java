import com.azure.resourcemanager.apimanagement.models.ApiType;
import com.azure.resourcemanager.apimanagement.models.Protocol;
import java.util.Arrays;

/** Samples for Api CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateWebsocketApi.json
     */
    /**
     * Sample code: ApiManagementCreateWebSocketApi.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateWebSocketApi(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apis()
            .define("tempgroup")
            .withExistingService("rg1", "apimService1")
            .withDisplayName("apiname1463")
            .withServiceUrl("wss://echo.websocket.org")
            .withPath("newapiPath")
            .withProtocols(Arrays.asList(Protocol.WSS, Protocol.WS))
            .withDescription("apidescription5200")
            .withApiType(ApiType.WEBSOCKET)
            .create();
    }
}
