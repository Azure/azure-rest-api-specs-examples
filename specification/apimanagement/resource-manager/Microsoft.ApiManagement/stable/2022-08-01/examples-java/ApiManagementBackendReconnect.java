import com.azure.resourcemanager.apimanagement.models.BackendReconnectContract;
import java.time.Duration;

/** Samples for Backend Reconnect. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementBackendReconnect.json
     */
    /**
     * Sample code: ApiManagementBackendReconnect.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementBackendReconnect(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .backends()
            .reconnectWithResponse(
                "rg1",
                "apimService1",
                "proxybackend",
                new BackendReconnectContract().withAfter(Duration.parse("PT3S")),
                com.azure.core.util.Context.NONE);
    }
}
