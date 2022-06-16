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
