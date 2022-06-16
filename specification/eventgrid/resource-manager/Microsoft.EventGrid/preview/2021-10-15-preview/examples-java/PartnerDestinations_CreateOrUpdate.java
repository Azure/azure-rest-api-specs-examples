import java.time.OffsetDateTime;
import java.util.UUID;

/** Samples for PartnerDestinations CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/PartnerDestinations_CreateOrUpdate.json
     */
    /**
     * Sample code: PartnerDestinations_CreateOrUpdate.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerDestinationsCreateOrUpdate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .partnerDestinations()
            .define("examplePartnerDestinationName1")
            .withRegion("westus2")
            .withExistingResourceGroup("examplerg")
            .withPartnerRegistrationImmutableId(UUID.fromString("0bd70ee2-7d95-447e-ab1f-c4f320019404"))
            .withEndpointServiceContext("This is an example")
            .withExpirationTimeIfNotActivatedUtc(OffsetDateTime.parse("2022-03-14T19:33:43.430Z"))
            .withEndpointBaseUrl("https://www.example/endpoint")
            .withMessageForActivation("Sample Activation message")
            .create();
    }
}
