Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-communication_1.1.0-beta.1/sdk/communication/azure-resourcemanager-communication/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for CommunicationService ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/stable/2020-08-20/examples/listKeys.json
     */
    /**
     * Sample code: List keys.
     *
     * @param manager Entry point to CommunicationManager.
     */
    public static void listKeys(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager
            .communicationServices()
            .listKeysWithResponse("MyResourceGroup", "MyCommunicationResource", Context.NONE);
    }
}
```
