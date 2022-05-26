Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-communication_1.1.0-beta.2/sdk/communication/azure-resourcemanager-communication/README.md) on how to add the SDK to your project and authenticate.

```java
/** Samples for EmailServices CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2021-10-01-preview/examples/emailServices/createOrUpdate.json
     */
    /**
     * Sample code: Create or update EmailService resource.
     *
     * @param manager Entry point to CommunicationManager.
     */
    public static void createOrUpdateEmailServiceResource(
        com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager
            .emailServices()
            .define("MyEmailServiceResource")
            .withRegion("Global")
            .withExistingResourceGroup("MyResourceGroup")
            .withDataLocation("United States")
            .create();
    }
}
```
