Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-eventgrid_1.2.0-beta.1/sdk/eventgrid/azure-resourcemanager-eventgrid/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.eventgrid.models.PartnerRegistration;

/** Samples for PartnerRegistrations Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/PartnerRegistrations_Update.json
     */
    /**
     * Sample code: PartnerRegistrations_Update.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerRegistrationsUpdate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        PartnerRegistration resource =
            manager
                .partnerRegistrations()
                .getByResourceGroupWithResponse("examplerg", "examplePartnerRegistrationName1", Context.NONE)
                .getValue();
        resource
            .update()
            .withSetupUri("https://www.example.com/newsetup.html")
            .withLogoUri("https://www.example.com/newlogo.png")
            .apply();
    }
}
```
