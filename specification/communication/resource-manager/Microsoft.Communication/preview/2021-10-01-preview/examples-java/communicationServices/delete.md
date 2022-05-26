Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-communication_1.1.0-beta.2/sdk/communication/azure-resourcemanager-communication/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for CommunicationServices Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2021-10-01-preview/examples/communicationServices/delete.json
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
