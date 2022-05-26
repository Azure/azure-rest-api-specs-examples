Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-communication_1.1.0-beta.2/sdk/communication/azure-resourcemanager-communication/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.communication.models.KeyType;
import com.azure.resourcemanager.communication.models.RegenerateKeyParameters;

/** Samples for CommunicationServices RegenerateKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2021-10-01-preview/examples/communicationServices/regenerateKey.json
     */
    /**
     * Sample code: Regenerate key.
     *
     * @param manager Entry point to CommunicationManager.
     */
    public static void regenerateKey(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager
            .communicationServices()
            .regenerateKey(
                "MyResourceGroup",
                "MyCommunicationResource",
                new RegenerateKeyParameters().withKeyType(KeyType.PRIMARY),
                Context.NONE);
    }
}
```
