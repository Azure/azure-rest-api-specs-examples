```java
/** Samples for CommunicationService CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/stable/2020-08-20/examples/createOrUpdate.json
     */
    /**
     * Sample code: Create or update resource.
     *
     * @param manager Entry point to CommunicationManager.
     */
    public static void createOrUpdateResource(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager
            .communicationServices()
            .define("MyCommunicationResource")
            .withExistingResourceGroup("MyResourceGroup")
            .withRegion("Global")
            .withDataLocation("United States")
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-communication_1.1.0-beta.1/sdk/communication/azure-resourcemanager-communication/README.md) on how to add the SDK to your project and authenticate.
