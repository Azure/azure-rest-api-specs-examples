```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.BackendReconnectContract;
import java.time.Duration;

/** Samples for Backend Reconnect. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementBackendReconnect.json
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
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
