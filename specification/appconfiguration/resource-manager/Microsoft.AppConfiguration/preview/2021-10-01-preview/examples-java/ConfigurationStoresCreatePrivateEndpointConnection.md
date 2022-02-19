Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-appconfiguration_1.0.0-beta.4/sdk/appconfiguration/azure-resourcemanager-appconfiguration/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.appconfiguration.models.ConnectionStatus;
import com.azure.resourcemanager.appconfiguration.models.PrivateLinkServiceConnectionState;

/** Samples for PrivateEndpointConnections CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/preview/2021-10-01-preview/examples/ConfigurationStoresCreatePrivateEndpointConnection.json
     */
    /**
     * Sample code: PrivateEndpointConnection_CreateOrUpdate.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void privateEndpointConnectionCreateOrUpdate(
        com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager
            .privateEndpointConnections()
            .define("myConnection")
            .withExistingConfigurationStore("myResourceGroup", "contoso")
            .withPrivateLinkServiceConnectionState(
                new PrivateLinkServiceConnectionState()
                    .withStatus(ConnectionStatus.APPROVED)
                    .withDescription("Auto-Approved"))
            .create();
    }
}
```
