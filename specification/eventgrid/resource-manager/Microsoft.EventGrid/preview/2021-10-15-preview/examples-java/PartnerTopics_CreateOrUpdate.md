```java
import java.time.OffsetDateTime;
import java.util.UUID;

/** Samples for PartnerTopics CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/PartnerTopics_CreateOrUpdate.json
     */
    /**
     * Sample code: PartnerTopics_CreateOrUpdate.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void partnerTopicsCreateOrUpdate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .partnerTopics()
            .define("examplePartnerTopicName1")
            .withRegion("westus2")
            .withExistingResourceGroup("examplerg")
            .withPartnerRegistrationImmutableId(UUID.fromString("6f541064-031d-4cc8-9ec3-a3b4fc0f7185"))
            .withSource("ContosoCorp.Accounts.User1")
            .withExpirationTimeIfNotActivatedUtc(OffsetDateTime.parse("2022-03-23T23:06:13.109Z"))
            .withPartnerTopicFriendlyDescription("Example description")
            .withMessageForActivation("Example message for activation")
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-eventgrid_1.2.0-beta.2/sdk/eventgrid/azure-resourcemanager-eventgrid/README.md) on how to add the SDK to your project and authenticate.
