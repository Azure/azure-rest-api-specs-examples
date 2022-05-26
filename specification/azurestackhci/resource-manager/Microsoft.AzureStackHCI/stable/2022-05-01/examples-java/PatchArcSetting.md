Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-azurestackhci_1.0.0-beta.3/sdk/azurestackhci/azure-resourcemanager-azurestackhci/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.azurestackhci.models.ArcSetting;
import java.io.IOException;

/** Samples for ArcSettings Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2022-05-01/examples/PatchArcSetting.json
     */
    /**
     * Sample code: Patch ArcSetting.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void patchArcSetting(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager)
        throws IOException {
        ArcSetting resource =
            manager.arcSettings().getWithResponse("test-rg", "myCluster", "default", Context.NONE).getValue();
        resource
            .update()
            .withConnectivityProperties(
                SerializerFactory
                    .createDefaultManagementSerializerAdapter()
                    .deserialize("{\"enabled\":true}", Object.class, SerializerEncoding.JSON))
            .apply();
    }
}
```
