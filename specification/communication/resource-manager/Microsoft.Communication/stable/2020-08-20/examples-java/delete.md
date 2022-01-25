Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-communication_1.1.0-beta.1/sdk/communication/azure-resourcemanager-communication/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for CommunicationService Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/stable/2020-08-20/examples/delete.json
     */
    /**
     * Sample code: Delete resource.
     *
     * @param manager Entry point to CommunicationManager.
     */
    public static void deleteResource(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.communicationServices().delete("MyResourceGroup", "MyCommunicationResource", Context.NONE);
    }
}
```
