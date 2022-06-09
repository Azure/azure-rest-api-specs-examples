```java
import com.azure.core.util.Context;

/** Samples for Domains Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2021-10-01-preview/examples/domains/delete.json
     */
    /**
     * Sample code: Delete Domains resource.
     *
     * @param manager Entry point to CommunicationManager.
     */
    public static void deleteDomainsResource(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.domains().delete("MyResourceGroup", "MyEmailServiceResource", "mydomain.com", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-communication_1.1.0-beta.2/sdk/communication/azure-resourcemanager-communication/README.md) on how to add the SDK to your project and authenticate.
