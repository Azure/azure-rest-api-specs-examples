Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-azurestackhci_1.0.0-beta.3/sdk/azurestackhci/azure-resourcemanager-azurestackhci/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import java.io.IOException;

/** Samples for Extensions Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2022-05-01/examples/PutExtension.json
     */
    /**
     * Sample code: Create Arc Extension.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void createArcExtension(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager)
        throws IOException {
        manager
            .extensions()
            .define("MicrosoftMonitoringAgent")
            .withExistingArcSetting("test-rg", "myCluster", "default")
            .withPublisher("Microsoft.Compute")
            .withTypePropertiesType("MicrosoftMonitoringAgent")
            .withTypeHandlerVersion("1.10")
            .withSettings(
                SerializerFactory
                    .createDefaultManagementSerializerAdapter()
                    .deserialize("{\"workspaceId\":\"xx\"}", Object.class, SerializerEncoding.JSON))
            .withProtectedSettings(
                SerializerFactory
                    .createDefaultManagementSerializerAdapter()
                    .deserialize("{\"workspaceKey\":\"xx\"}", Object.class, SerializerEncoding.JSON))
            .create();
    }
}
```
