```java
import com.azure.resourcemanager.apimanagement.models.PrivateEndpointConnectionRequestProperties;
import com.azure.resourcemanager.apimanagement.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.apimanagement.models.PrivateLinkServiceConnectionState;

/** Samples for PrivateEndpointConnection CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementApproveOrRejectPrivateEndpointConnection.json
     */
    /**
     * Sample code: ApiManagementApproveOrRejectPrivateEndpointConnection.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementApproveOrRejectPrivateEndpointConnection(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .privateEndpointConnections()
            .define("privateEndpointConnectionName")
            .withExistingService("rg1", "apimService1")
            .withProperties(
                new PrivateEndpointConnectionRequestProperties()
                    .withPrivateLinkServiceConnectionState(
                        new PrivateLinkServiceConnectionState()
                            .withStatus(PrivateEndpointServiceConnectionStatus.APPROVED)
                            .withDescription("The Private Endpoint Connection is approved.")))
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
