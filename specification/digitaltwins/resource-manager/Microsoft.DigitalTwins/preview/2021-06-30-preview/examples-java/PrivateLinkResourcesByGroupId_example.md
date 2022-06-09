```java
import com.azure.core.util.Context;

/** Samples for PrivateLinkResources Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/preview/2021-06-30-preview/examples/PrivateLinkResourcesByGroupId_example.json
     */
    /**
     * Sample code: Get the specified private link resource for the given Digital Twin.
     *
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void getTheSpecifiedPrivateLinkResourceForTheGivenDigitalTwin(
        com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager.privateLinkResources().getWithResponse("resRg", "myDigitalTwinsService", "subResource", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-digitaltwins_1.0.0-beta.2/sdk/digitaltwins/azure-resourcemanager-digitaltwins/README.md) on how to add the SDK to your project and authenticate.
