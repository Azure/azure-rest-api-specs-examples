
import com.azure.resourcemanager.eventgrid.models.ChannelType;
import com.azure.resourcemanager.eventgrid.models.PartnerTopicInfo;
import java.time.OffsetDateTime;

/**
 * Samples for Channels CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/Channels_CreateOrUpdate.
     * json
     */
    /**
     * Sample code: Channels_CreateOrUpdate.
     * 
     * @param manager Entry point to EventGridManager.
     */
    public static void channelsCreateOrUpdate(com.azure.resourcemanager.eventgrid.EventGridManager manager) {
        manager.channels().define("exampleChannelName1")
            .withExistingPartnerNamespace("examplerg", "examplePartnerNamespaceName1")
            .withChannelType(ChannelType.PARTNER_TOPIC)
            .withPartnerTopicInfo(new PartnerTopicInfo().withAzureSubscriptionId("5b4b650e-28b9-4790-b3ab-ddbd88d727c4")
                .withResourceGroupName("examplerg2").withName("examplePartnerTopic1")
                .withSource("ContosoCorp.Accounts.User1"))
            .withMessageForActivation("Example message to approver")
            .withExpirationTimeIfNotActivatedUtc(OffsetDateTime.parse("2021-10-21T22:50:25.410433Z")).create();
    }
}
