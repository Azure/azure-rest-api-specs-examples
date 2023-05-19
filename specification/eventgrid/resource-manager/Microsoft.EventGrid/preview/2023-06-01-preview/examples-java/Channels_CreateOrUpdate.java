import com.azure.resourcemanager.eventgrid.models.ChannelType;
import com.azure.resourcemanager.eventgrid.models.PartnerTopicInfo;
import java.time.OffsetDateTime;

/** Samples for Channels CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/Channels_CreateOrUpdate.json
     */
    /**
     * Sample code: Channels_CreateOrUpdate.
     *
     * @param manager Entry point to EventGridManager.
     */
    public static void channelsCreateOrUpdate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager
            .channels()
            .define("exampleChannelName1")
            .withExistingPartnerNamespace("examplerg", "examplePartnerNamespaceName1")
            .withChannelType(ChannelType.PARTNER_TOPIC)
            .withPartnerTopicInfo(
                new PartnerTopicInfo()
                    .withAzureSubscriptionId("8f6b6269-84f2-4d09-9e31-1127efcd1e40")
                    .withResourceGroupName("examplerg2")
                    .withName("examplePartnerTopic1")
                    .withSource("ContosoCorp.Accounts.User1"))
            .withMessageForActivation("Example message to approver")
            .withExpirationTimeIfNotActivatedUtc(OffsetDateTime.parse("2021-10-21T22:50:25.410433Z"))
            .create();
    }
}
