Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-communication_1.1.0-beta.1/sdk/communication/azure-resourcemanager-communication/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.communication.models.NameAvailabilityParameters;

/** Samples for CommunicationService CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/stable/2020-08-20/examples/checkNameAvailabilityAvailable.json
     */
    /**
     * Sample code: Check name availability available.
     *
     * @param manager Entry point to CommunicationManager.
     */
    public static void checkNameAvailabilityAvailable(
        com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager
            .communicationServices()
            .checkNameAvailabilityWithResponse(
                new NameAvailabilityParameters()
                    .withType("Microsoft.Communication/CommunicationServices")
                    .withName("MyCommunicationService"),
                Context.NONE);
    }
}
```
