Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-digitaltwins_1.0.0-beta.2/sdk/digitaltwins/azure-resourcemanager-digitaltwins/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.digitaltwins.models.AuthenticationType;
import com.azure.resourcemanager.digitaltwins.models.ServiceBus;

/** Samples for DigitalTwinsEndpoint CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/preview/2021-06-30-preview/examples/DigitalTwinsEndpointPut_example.json
     */
    /**
     * Sample code: Put a DigitalTwinsEndpoint resource.
     *
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void putADigitalTwinsEndpointResource(
        com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager
            .digitalTwinsEndpoints()
            .define("myServiceBus")
            .withExistingDigitalTwinsInstance("resRg", "myDigitalTwinsService")
            .withProperties(
                new ServiceBus()
                    .withAuthenticationType(AuthenticationType.KEY_BASED)
                    .withPrimaryConnectionString(
                        "Endpoint=sb://mysb.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=xyzxyzoX4=;EntityPath=abcabc")
                    .withSecondaryConnectionString(
                        "Endpoint=sb://mysb.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=xyzxyzoX4=;EntityPath=abcabc"))
            .create();
    }
}
```
