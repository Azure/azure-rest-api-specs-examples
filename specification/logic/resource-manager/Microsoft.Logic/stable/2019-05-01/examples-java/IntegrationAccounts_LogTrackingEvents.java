
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.logic.models.EventLevel;
import com.azure.resourcemanager.logic.models.TrackingEvent;
import com.azure.resourcemanager.logic.models.TrackingEventErrorInfo;
import com.azure.resourcemanager.logic.models.TrackingEventsDefinition;
import com.azure.resourcemanager.logic.models.TrackingRecordType;
import java.io.IOException;
import java.time.OffsetDateTime;
import java.util.Arrays;

/**
 * Samples for IntegrationAccounts LogTrackingEvents.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/
     * IntegrationAccounts_LogTrackingEvents.json
     */
    /**
     * Sample code: Log a tracked event.
     * 
     * @param manager Entry point to LogicManager.
     */
    public static void logATrackedEvent(com.azure.resourcemanager.logic.LogicManager manager) throws IOException {
        manager.integrationAccounts().logTrackingEventsWithResponse("testResourceGroup", "testIntegrationAccount",
            new TrackingEventsDefinition().withSourceType("Microsoft.Logic/workflows")
                .withEvents(Arrays.asList(new TrackingEvent().withEventLevel(EventLevel.INFORMATIONAL)
                    .withEventTime(OffsetDateTime.parse("2016-08-05T01:54:49.505567Z"))
                    .withRecordType(TrackingRecordType.AS2MESSAGE)
                    .withRecord(SerializerFactory.createDefaultManagementSerializerAdapter().deserialize(
                        "{\"agreementProperties\":{\"agreementName\":\"testAgreement\",\"as2From\":\"testas2from\",\"as2To\":\"testas2to\",\"receiverPartnerName\":\"testPartner2\",\"senderPartnerName\":\"testPartner1\"},\"messageProperties\":{\"IsMessageEncrypted\":false,\"IsMessageSigned\":false,\"correlationMessageId\":\"Unique message identifier\",\"direction\":\"Receive\",\"dispositionType\":\"received-success\",\"fileName\":\"test\",\"isMdnExpected\":true,\"isMessageCompressed\":false,\"isMessageFailed\":false,\"isNrrEnabled\":true,\"mdnType\":\"Async\",\"messageId\":\"12345\"}}",
                        Object.class, SerializerEncoding.JSON))
                    .withError(new TrackingEventErrorInfo().withMessage("Some error occurred")
                        .withCode("fakeTokenPlaceholder")))),
            com.azure.core.util.Context.NONE);
    }
}
